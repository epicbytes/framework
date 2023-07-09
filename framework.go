package framework

import (
	"context"
	"errors"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/epicbytes/framework/bus"
	"github.com/epicbytes/framework/config"
	mqtt3 "github.com/epicbytes/framework/pubsub/mqtt"
	redis2 "github.com/epicbytes/framework/pubsub/redis"
	"github.com/epicbytes/framework/runtime"
	"github.com/epicbytes/framework/s3"
	"github.com/epicbytes/framework/service"
	"github.com/epicbytes/framework/storage/mongodb"
	"github.com/epicbytes/framework/tasks"
	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"net/http"
	"sync"
	"time"
)

const (
	listenerReplicationStarted = "mongoReplicationStarted"
	listenerReplication        = "mongoReplication"
)

type frmwrk struct {
	ctx                 context.Context
	OnStartup           func(ctx context.Context)
	OnFinish            func(ctx context.Context)
	config              *config.Config
	databaseClient      *mongo.Client
	Models              sync.Map
	Stores              map[mongodb.CollectionName]*sync.Map
	TaskManagers        sync.Map
	Workers             sync.Map
	GRPCClients         sync.Map
	InternalGRPCClients sync.Map
	PubSub              *redis.Client
	MqttClient          mqtt.Client
	Gateway             BaseGateway
	grpcRoutes          *http.ServeMux
	grpcInternalRoutes  *http.ServeMux
	TGBot               *tgbotapi.BotAPI
	S3                  s3.MinioStorageInt
}

type Framework interface {
	Run() error
	SetOnStartup(fn func(ctx context.Context))
	SetOnFinish(fn func(ctx context.Context))
	OnMongoReplicationModelStart(fn func(ctx context.Context, data mongodb.ReplicationEvent))
	OnMongoReplication(fn func(ctx context.Context, data mongodb.ReplicationEvent))
	RegisterGRPCServer(server http.Handler, path string, middlewares ...func(handler http.Handler) http.Handler)
	RegisterInternalGRPCServer(server http.Handler, path string, middlewares ...func(handler http.Handler) http.Handler)
	GetContext() context.Context
	GetConfig() *config.Config
	GetModel(name mongodb.CollectionName) mongodb.Model
	GetStore(name mongodb.CollectionName) *sync.Map
	GetTaskManager(name string) (client.Client, error)
	RegisterWorker(namespace string, queue string, options worker.Options, onCreate func(wrk worker.Worker)) error
	RegisterGRPCClient(clientName string, client interface{})
	RegisterInternalGRPCClient(clientName string, client interface{})
	GetGRPCClient(name string) interface{}
	GetInternalGRPCClient(name string) interface{}
	RegisterNetListener(server BaseGateway)
	GetMQTTClient() mqtt.Client
	GetTGBot() *tgbotapi.BotAPI
	MQTTPublish(topic string, obj interface{})
	GetS3Client() s3.MinioStorageInt
}

type ForGateway interface {
	GetContext() context.Context
	GetConfig() *config.Config
	GetModel(name mongodb.CollectionName) mongodb.Model
	GetTaskManager(name string) client.Client
}

func (f *frmwrk) SetOnStartup(fn func(ctx context.Context)) {
	f.OnStartup = fn
}

func (f *frmwrk) SetOnFinish(fn func(ctx context.Context)) {
	f.OnFinish = fn
}

func (f *frmwrk) OnMongoReplicationModelStart(fn func(ctx context.Context, data mongodb.ReplicationEvent)) {
	bus.Listen(listenerReplicationStarted, func(message bus.Message) {
		data := message.(mongodb.ReplicationEvent)
		fn(f.ctx, data)
	})
}

func (f *frmwrk) OnMongoReplication(fn func(ctx context.Context, replicationEvent mongodb.ReplicationEvent)) {
	bus.Listen(listenerReplication, func(message bus.Message) {
		data := message.(mongodb.ReplicationEvent)
		fn(f.ctx, data)
	})
}

func (f *frmwrk) RegisterGRPCServer(server http.Handler, path string, middlewares ...func(handler http.Handler) http.Handler) {
	if f.grpcRoutes == nil {
		f.grpcRoutes = http.NewServeMux()
	}
	for _, middleware := range middlewares {
		server = middleware(server)
	}
	f.grpcRoutes.Handle(path, server)
}

func (f *frmwrk) RegisterInternalGRPCServer(server http.Handler, path string, middlewares ...func(handler http.Handler) http.Handler) {
	if f.grpcInternalRoutes == nil {
		f.grpcInternalRoutes = http.NewServeMux()
	}
	hnd := server
	for _, middleware := range middlewares {
		hnd = middleware(hnd)
	}
	f.grpcInternalRoutes.Handle(path, hnd)
}

func (f *frmwrk) RegisterNetListener(server BaseGateway) {
	f.Gateway = server
}

func (f *frmwrk) RegisterGRPCClient(clientName string, client interface{}) {
	f.GRPCClients.Store(clientName, client)
}

func (f *frmwrk) RegisterInternalGRPCClient(clientName string, client interface{}) {
	f.InternalGRPCClients.Store(clientName, client)
}

func (f *frmwrk) RegisterWorker(namespace string, queue string, options worker.Options, fn func(wrk worker.Worker)) error {
	cl, err := f.GetTaskManager(namespace)
	if err != nil {
		return err
	}
	wrk := worker.New(cl, queue, options)
	fn(wrk)
	f.Workers.Store(namespace+"-"+queue, wrk)
	return nil
}

func (f *frmwrk) GetContext() context.Context {
	return f.ctx
}

func (f *frmwrk) GetMQTTClient() mqtt.Client {
	return f.MqttClient
}

func (f *frmwrk) MQTTPublish(topic string, obj interface{}) {
	if f.MqttClient != nil {
		f.MqttClient.Publish(topic, 0, false, mqtt3.ConvertForMQTT(obj))
	}
}

func (f *frmwrk) GetConfig() *config.Config {
	return f.config
}

func (f *frmwrk) GetModel(name mongodb.CollectionName) mongodb.Model {
	model, ok := f.Models.Load(name.String())
	if !ok {
		return nil
	}
	return model.(mongodb.Model)
}

func (f *frmwrk) GetGRPCClient(name string) interface{} {
	grpcClient, ok := f.GRPCClients.Load(name)
	if !ok {
		log.Error().Err(errors.New("can`t load grpc client: " + name))
		return nil
	}
	return grpcClient
}

func (f *frmwrk) GetInternalGRPCClient(name string) interface{} {
	internalGRPCClient, ok := f.InternalGRPCClients.Load(name)
	if !ok {
		log.Error().Err(errors.New("can`t load grpc client: " + name))
		return nil
	}
	return internalGRPCClient
}

func (f *frmwrk) GetStore(name mongodb.CollectionName) *sync.Map {
	store, ok := f.Stores[name]
	if !ok {
		return nil
	}
	return store
}

func (f *frmwrk) GetTaskManager(name string) (client.Client, error) {
	model, ok := f.TaskManagers.Load(name)
	if !ok {
		return nil, errors.New("task manager is not connected")
	}
	return model.(client.Client), nil
}

func (f *frmwrk) GetTGBot() *tgbotapi.BotAPI {
	return f.TGBot
}

func (f *frmwrk) GetS3Client() s3.MinioStorageInt {
	return f.S3
}

// Run main livecycle
func (f *frmwrk) Run() error {

	var tasksRuntime []runtime.Service

	/*
		if f.Gateway != nil {
			f.Gateway.SetFramework(f)
			wg.Go(func() error {
				if f.config.Gateway.Addr == "" {
					log.Info().Msg("Stack started")
				} else {
					log.Info().Msgf("Gateway started at %s\n", f.config.Gateway.Addr)
				}
				f.Gateway.Start(f.ctx)
				return nil
			})

		}

		f.Workers.Range(func(key, value any) bool {
			go func() {
				err := value.(worker.Worker).Run(worker.InterruptCh())
				log.Error().Err(err)
				//return err
			}()
			return true
		})*/

	if len(f.config.Server.Addr) > 0 {
		tasksRuntime = append(tasksRuntime, &service.GRPCService{
			GrpcMultiplexer: f.grpcRoutes,
			Config:          f.config,
		})
	}

	if len(f.config.Server.InternalAddr) > 0 {
		tasksRuntime = append(tasksRuntime, &service.GRPCInternalService{
			GrpcInternalMultiplexer: f.grpcInternalRoutes,
			Config:                  f.config,
		})
	}

	if len(f.config.Mongo.URI) > 0 {
		mongoClient := &mongodb.MongoClient{
			URI: f.config.Mongo.URI,
		}
		mongoClient.OnConnect(func(ctx context.Context, client *mongo.Client) error {
			for _, entity := range f.config.Mongo.Entities {
				f.Models.Store(entity.Collection.String(), mongodb.New(client, f.config.Mongo.DatabaseName, entity.Collection))
				if entity.WithMemstore {
					f.Stores[entity.Collection] = &sync.Map{}
				}
				if entity.WithReplication {
					mdl, ok := f.Models.Load(entity.Collection.String())
					if !ok {
						return errors.New("replication: model is not set")
					}
					log.Debug().Msgf("Start replication for %s", entity.Collection.String())
					cursor, err := mdl.(mongodb.Model).GetCollection().Find(f.ctx, bson.M{})
					if err != nil {
						return err
					}
					var results []bson.M
					if err = cursor.All(context.TODO(), &results); err != nil {
						log.Error().Err(err)
					}
					for _, result := range results {
						id := result["_id"].(primitive.ObjectID).Hex()
						data, _ := bson.Marshal(result)
						bus.Ring(listenerReplicationStarted, mongodb.ReplicationEvent{
							Id:             id,
							Type:           mongodb.ReplicationEventInsert,
							CollectionName: entity.Collection,
							Data:           data,
						})
					}

					go (func() {
						cs, err := mdl.(mongodb.Model).GetCollection().Watch(context.TODO(), entity.ReplicationQuery)
						if err != nil {
							log.Error().Err(err)
							panic(err)
						}
						defer cs.Close(context.TODO())
						for cs.Next(context.TODO()) {
							event, err := mongodb.WatchEventHandler(cs)
							if err != nil {
								log.Error().Err(err)
								return
							}

							if err := bus.Ring(listenerReplication, mongodb.ReplicationEvent{
								Id:             event.ID,
								Type:           event.OperationType,
								CollectionName: entity.Collection,
								Data:           event.Data,
							}); err != nil {
								log.Error().Err(err)
								return
							}

							log.Printf("Event: %s, %s", event.ID, event.Collection)
						}
						if err := cs.Err(); err != nil {
							log.Error().Err(err)
							return
						}
						return
					})()
				}
			}
			return nil
		})
		tasksRuntime = append(tasksRuntime, mongoClient)
	}

	if len(f.config.Redis.URI) > 0 {
		redisClient := &redis2.RedisClient{
			URI:      f.config.Redis.URI,
			Password: f.config.Redis.Password,
		}
		redisClient.OnConnect(func(ctx context.Context, client *redis.Client) error {
			f.PubSub = client
			return nil
		})
		tasksRuntime = append(tasksRuntime, redisClient)
	}

	if len(f.config.S3.Address) > 0 {
		s3Client := &s3.MinioStorage{Config: f.config}
		f.S3 = s3Client
		tasksRuntime = append(tasksRuntime, s3Client)
	}

	if len(f.config.MQTTClient.URI) > 0 {
		mqttClient := &mqtt3.MQTTClient{
			Config: f.config,
		}
		mqttClient.OnConnect(func(ctx context.Context, client mqtt.Client) error {
			f.MqttClient = client
			return nil
		})
		tasksRuntime = append(tasksRuntime, mqttClient)
	}

	var keeper = runtime.TaskKeeper{
		Tasks:           tasksRuntime,
		ShutdownTimeout: time.Second * 10,
		PingPeriod:      time.Millisecond * 500,
	}
	var app = runtime.Application{
		MainFunc: func(ctx context.Context, halt <-chan struct{}) error {
			var errShutdown = make(chan error, 1)
			f.OnStartup(f.ctx)
			defer f.OnFinish(f.ctx)
			go func() {
				defer close(errShutdown)
				select {
				case <-halt:
				case <-ctx.Done():

				}
			}()
			err, ok := <-errShutdown
			if ok {
				return err
			}
			return nil
		},
		Resources:          &keeper,
		TerminationTimeout: time.Second * 10,
	}
	return app.Run()
}

func New(cfg *config.Config) (framework Framework) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	var (
		ctx = context.Background()
		frm = &frmwrk{
			ctx:    ctx,
			config: cfg,
			Stores: map[mongodb.CollectionName]*sync.Map{},
		}
	)

	if frm.config.Temporal.URI != "" {
		for _, ns := range frm.config.Temporal.Namespaces {
			tm, err := tasks.New(frm.config.Temporal.URI, ns)
			if err != nil {
				log.Error().Err(err)
			}
			frm.TaskManagers.Store(ns, tm)
		}
	}

	if frm.config.Telegram.APIToken != "" {
		tbot, err := tgbotapi.NewBotAPI(frm.config.Telegram.APIToken)
		if err != nil {
			panic(err)
		}

		frm.TGBot = tbot

		frm.TGBot.Debug = true

	}

	return frm
}
