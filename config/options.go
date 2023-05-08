package config

import (
	"github.com/epicbytes/framework/storage/mongodb"
	"github.com/panjf2000/gnet/v2"
)

type (
	Config struct {
		Server     serverOption   `envPrefix:"SERVER_"`
		Mongo      mongoOption    `envPrefix:"MONGO_"`
		Redis      redisOption    `envPrefix:"REDIS_"`
		MQTTClient mqttOption     `envPrefix:"MQTT_"`
		Temporal   temporalOption `envPrefix:"TEMPORAL_"`
		Gateway    gatewayOption  `envPrefix:"GATEWAY_"`
		Telegram   telegramOption `envPrefix:"TELEGRAM_"`
	}

	Option interface {
		apply(*Config)
	}

	optionsFromEnv Config

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
)

func (o optionsFromEnv) apply(opts *Config) {
	opts.Mongo = o.Mongo
	opts.Redis = o.Redis
	opts.Temporal = o.Temporal
}

func (o serverOption) apply(opts *Config) {
	opts.Server = o
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

func WithEnvFile(envfile string) Option {
	return optionsFromEnv{}
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
