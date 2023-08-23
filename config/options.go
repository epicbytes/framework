package config

import (
	"github.com/epicbytes/framework/storage/mongodb"
	"github.com/panjf2000/gnet/v2"
)

type (
	Config struct {
		HttpServer httpServerOption `envPrefix:"HTTP_SERVER_"`
		Server     serverOption     `envPrefix:"SERVER_"`
		Mongo      mongoOption      `envPrefix:"MONGO_"`
		Redis      redisOption      `envPrefix:"REDIS_"`
		MQTTClient mqttOption       `envPrefix:"MQTT_"`
		Temporal   temporalOption   `envPrefix:"TEMPORAL_"`
		Gateway    gatewayOption    `envPrefix:"GATEWAY_"`
		Telegram   telegramOption   `envPrefix:"TELEGRAM_"`
		S3         s3Option         `envPrefix:"S3_"`
	}

	Option interface {
		apply(*Config)
	}

	optionsFromEnv Config

	httpServerOption struct {
		Addr string `env:"ADDR"`
	}

	serverOption struct {
		Addr         string `env:"ADDR"`
		InternalAddr string `env:"INTERNAL_ADDR"`
	}

	mongoOption struct {
		URI          string `env:"URI"`
		DatabaseName string `env:"DATABASE_NAME"`
		Entities     []mongodb.ModelEntity
	}

	redisOption struct {
		URI      string `env:"URI"`
		Password string `env:"PASSWORD"`
		Database int    `env:"DATABASE"`
	}

	mqttOption struct {
		URI      string `env:"URI"`
		Password string `env:"PASSWORD"`
		Username string `env:"USERNAME"`
		ClientId string `env:"CLIENT_ID"`
	}

	temporalOption struct {
		URI        string   `env:"URI"`
		Namespaces []string `env:"NAMESPACES" envSeparator:":"`
	}

	gatewayOption struct {
		Addr     string        `env:"ADDRESS"`
		Protocol string        `env:"PROTOCOL"`
		Options  []gnet.Option //todo: create setter
	}

	telegramOption struct {
		APIToken string `env:"API_TOKEN"`
	}

	s3Option struct {
		Address   string `env:"ADDRESS"`
		AccessKey string `env:"ACCESS_KEY"`
		SecretKey string `env:"SECRET_KEY"`
		Bucket    string `env:"BUCKET"`
		Region    string `env:"REGION"`
		Secure    bool   `env:"SECURE"`
	}
)

func (o optionsFromEnv) apply(opts *Config) {
	opts.Mongo = o.Mongo
	opts.Redis = o.Redis
	opts.Temporal = o.Temporal
	opts.S3 = o.S3
	opts.Gateway = o.Gateway
	opts.Server = o.Server
}

func (o serverOption) apply(opts *Config) {
	opts.Server = o
}

func (o httpServerOption) apply(opts *Config) {
	opts.HttpServer = o
}

func (o mongoOption) apply(opts *Config) {
	opts.Mongo = o
}

func (o redisOption) apply(opts *Config) {
	opts.Redis = o
}

func (o mqttOption) apply(opts *Config) {
	opts.MQTTClient = o
}

func (o temporalOption) apply(opts *Config) {
	opts.Temporal = o
}

func (o gatewayOption) apply(opts *Config) {
	opts.Gateway = o
}

func (o telegramOption) apply(opts *Config) {
	opts.Telegram = o
}

func (o s3Option) apply(opts *Config) {
	opts.S3 = o
}

func WithEnvFile(envfile string) Option {
	return optionsFromEnv{}
}

func WithHTTPServer(address string) Option {
	return httpServerOption{
		Addr: address,
	}
}

func WithGRPCServer(address string, internalAddress string) Option {
	return serverOption{
		Addr:         address,
		InternalAddr: internalAddress,
	}
}

func WithMongo(uriData string, databaseName string, entities ...mongodb.ModelEntity) Option {
	return mongoOption{
		URI:          uriData,
		DatabaseName: databaseName,
		Entities:     entities,
	}
}

func WithS3(Address, AccessKey, SecretKey, Bucket, Region string, Secure bool) Option {
	return s3Option{
		Address:   Address,
		AccessKey: AccessKey,
		SecretKey: SecretKey,
		Bucket:    Bucket,
		Region:    Region,
		Secure:    Secure,
	}
}

func (c *Config) SetMongoModels(entities ...mongodb.ModelEntity) {
	c.Mongo.Entities = entities
}

func WithRedis(uriData string, password string, database int) Option {
	return redisOption{
		URI:      uriData,
		Password: password,
		Database: database,
	}
}

func WithMQTT(uriData string, password string, username string, clientID string) Option {
	return mqttOption{
		URI:      uriData,
		Password: password,
		Username: username,
		ClientId: clientID,
	}
}

func WithTemporal(uriData string, namespaces ...string) Option {
	return temporalOption{
		URI:        uriData,
		Namespaces: namespaces,
	}
}

func WithGateway(address string, protocol string, options ...gnet.Option) Option {
	return gatewayOption{
		Addr:     address,
		Protocol: protocol,
		Options:  options,
	}
}

func WithTelegram(apikey string) Option {
	return telegramOption{
		APIToken: apikey,
	}
}
