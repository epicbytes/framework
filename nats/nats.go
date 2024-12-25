package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func newJetStream(cfg *Config) (jetstream.JetStream, error) {
	nc, err := nats.Connect(cfg.Endpoint, nats.UserInfo(cfg.User, cfg.Pass))
	if err != nil {
		return nil, err
	}

	return jetstream.New(nc)
}
