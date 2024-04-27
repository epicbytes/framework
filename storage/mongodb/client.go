package mongodb

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	ctx       context.Context
	client    *mongo.Client
	URI       string
	onConnect func(ctx context.Context, client *mongo.Client) error
}

func (t *MongoClient) OnConnect(fn func(ctx context.Context, client *mongo.Client) error) {
	t.onConnect = fn
}

func (t *MongoClient) Init(ctx context.Context) error {
	t.ctx = ctx
	var err error
	log.Debug().Msg("INITIAL MongoDB")
	t.client, err = mongo.Connect(ctx, options2.Client().ApplyURI(t.URI) /*SetMonitor(cmdMonitor)*/)
	if err != nil {
		return err
	}
	if t.onConnect != nil {
		err = t.onConnect(t.ctx, t.client)
		if err != nil {
			log.Error().Err(err).Send()
			return err
		}
	}

	return nil
}

func (t *MongoClient) GetClient() *mongo.Client {
	return t.client
}

func (t *MongoClient) Ping(context.Context) error {
	return nil
}

func (t *MongoClient) Close() error {
	log.Debug().Msg("CLOSE MongoDB connection")
	return t.client.Disconnect(t.ctx)
}
