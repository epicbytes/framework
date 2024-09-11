package kafka

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"time"
)

type kafka struct {
	producer sarama.AsyncProducer
	consumer sarama.Consumer
	logger   *zap.Logger
	Done     chan struct{}
}

type Kafka interface {
	Produce(message *sarama.ProducerMessage)
	Consume(topic string, fn func(ctx context.Context, msg *sarama.ConsumerMessage), opts ...ConsumerOption)
}

type KafkaTech interface {
	StartKafka(ctx context.Context) error
	StopKafka(ctx context.Context) error
	Logger() *zap.Logger
}

func (k *kafka) Produce(message *sarama.ProducerMessage) {
	k.producer.Input() <- message
}

func (k *kafka) Consume(topic string, fn func(ctx context.Context, msg *sarama.ConsumerMessage), opts ...ConsumerOption) {

	options := &consumerOption{
		Offset: sarama.OffsetOldest,
	}

	for _, opt := range opts {
		opt(options)
	}

	ctx := context.Background()
	partitionConsumer, err := k.consumer.ConsumePartition(topic, options.Partition, options.Offset)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			k.logger.Fatal("partition consumer error", zap.Error(err))
		}
	}()

ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			if options.timeout > 0 {
				ctxWithTimout, cancel := context.WithTimeout(ctx, options.timeout)
				fn(ctxWithTimout, msg)
				cancel()
			} else {
				fn(ctx, msg)
			}
		case <-k.Done:
			k.logger.Warn("consumer stopped listen events for topic", zap.String("topic", topic))
			break ConsumerLoop
		}
	}

}

func newKafka(logger *zap.Logger, cfg *Config) *kafka {
	version, err := cfg.ParsedVersion()
	if err != nil {
		logger.Fatal("Error parsing Kafka version", zap.Error(err))
	}

	sarama.Logger = newSaramaZapLogger(logger)

	config := sarama.NewConfig()
	config.Version = version
	config.ClientID = uuid.Must(uuid.NewV7()).String()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	//config.Producer.Return.Successes = true
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond

	if cfg.IsAuthRequired() {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = cfg.Username
		config.Net.SASL.Password = cfg.Password
		config.Net.SASL.Handshake = true
	}

	if cfg.IsCertificateAuth() {
		err := genTLSConfig(config, cfg.CertificateCrtPath, cfg.CertificateKeyPath, cfg.CertificateCAPath)
		if err != nil {
			logger.Fatal("Cert error", zap.Error(err))
		}
	}

	producer, err := sarama.NewAsyncProducer(cfg.Addrs, config)
	if err != nil {
		logger.Fatal("Failed to start Sarama producer:", zap.Error(err))
	}

	consumer, err := sarama.NewConsumer(cfg.Addrs, config)
	if err != nil {
		logger.Fatal("Failed to start Sarama consumer:", zap.Error(err))
	}

	return &kafka{
		logger:   logger,
		producer: producer,
		consumer: consumer,
		Done:     make(chan struct{}),
	}
}

func (k *kafka) StartKafka(ctx context.Context) error {
	go func() {
		for err := range k.producer.Errors() {
			k.logger.Warn("Failed to write access log entry:", zap.Error(err))
		}
	}()

	return nil
}

func (k *kafka) StopKafka(ctx context.Context) error {
	<-k.Done
	wg := &errgroup.Group{}
	wg.Go(func() error {
		return k.producer.Close()
	})
	wg.Go(func() error {
		return k.consumer.Close()
	})
	return wg.Wait()
}

func (k *kafka) Logger() *zap.Logger {
	return k.logger
}

func newSaramaZapLogger(logger *zap.Logger) sarama.StdLogger {
	sl, _ := zap.NewStdLogAt(logger, zapcore.DebugLevel)
	return sl
}

func genTLSConfig(config *sarama.Config, accessCertPath, accessKeyPath, caCertPath string) error {
	keypair, err := tls.LoadX509KeyPair(accessCertPath, accessKeyPath)
	if err != nil {
		return err
	}
	caCert, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		return err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{keypair},
		RootCAs:      caCertPool,
	}

	config.Net.TLS.Enable = true
	config.Net.TLS.Config = tlsConfig

	return nil
}
