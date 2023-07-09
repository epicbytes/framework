package mqtt

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/epicbytes/framework/config"
	"github.com/rs/zerolog/log"
)

type MQTTClient struct {
	ctx       context.Context
	client    mqtt.Client
	Config    *config.Config
	onConnect func(ctx context.Context, client mqtt.Client) error
}

func (t *MQTTClient) OnConnect(fn func(ctx context.Context, client mqtt.Client) error) {
	t.onConnect = fn
}

func (t *MQTTClient) Init(ctx context.Context) error {
	t.ctx = ctx
	var err error
	log.Debug().Msg("INITIAL MQTT")
	mqtOpt := mqtt.NewClientOptions()
	mqtOpt.AddBroker("tcp://" + t.Config.MQTTClient.URI)
	mqtOpt.SetUsername(t.Config.MQTTClient.Username)
	mqtOpt.SetPassword(t.Config.MQTTClient.Password)
	mqtOpt.SetClientID(t.Config.MQTTClient.ClientId)
	t.client = mqtt.NewClient(mqtOpt)

	if token := t.client.Connect(); token.Wait() && token.Error() != nil {
		log.Error().Err(token.Error())
		return err
	}

	if t.onConnect != nil {
		err = t.onConnect(t.ctx, t.client)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *MQTTClient) GetClient() mqtt.Client {
	return t.client
}

func (t *MQTTClient) Ping(context.Context) error {
	return nil
}

func (t *MQTTClient) Close() error {
	log.Debug().Msg("CLOSE MQTT connection")
	t.client.Disconnect(0)
	return nil
}
