package framework

import (
	"context"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/epicbytes/framework/bus"
	"github.com/epicbytes/framework/config"
	mqtt2 "github.com/epicbytes/framework/mqtt"
	redisPS "github.com/epicbytes/framework/pubsub/redis"
	"github.com/epicbytes/framework/storage/mongodb"
	"github.com/epicbytes/framework/tasks"
	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"sync"
)

const (
	listenerReplicationStarted = "mongoReplicationStarted"
	listenerReplication        = "mongoReplication"
)

type frmwrk struct {
	ctx                 context.Context
	onStartup           func(ctx context.Context)
	onShutdown          func(ctx context.Context)
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
	internalGrpcRoutes  *http.ServeMux
	TGBot               *tgbotapi.BotAPI
}

type Framework interface {
	Run() error
	OnStartup(fn func(ctx context.Context))
	OnFinish(fn func(ctx context.Context))
	OnMongoReplicationModelStart(fn func(ctx context.Context, data mongodb.ReplicationEvent))
	OnMongoReplication(fn func(ctx context.Context, data mongodb.ReplicationEvent))
	RegisterGRPCServer(server BaseTwirpServer, path string, middlewares ...func(handler http.Handler) http.Handler)
	RegisterInternalGRPCServer(server BaseTwirpServer, path string, middlewares ...func(handler http.Handler) http.Handler)
	GetContext() context.Context
	GetConfig() *config.Config
	GetModel(name mongodb.CollectionName) mongodb.Model
	GetStore(name mongodb.CollectionName) *sync.Map
	GetTaskManager(name string) client.Client
	RegisterWorker(namespace string, queue string, options worker.Options, onCreate func(wrk worker.Worker))
	RegisterGRPCClient(clientName string, client interface{})
	RegisterInternalGRPCClient(clientName string, client interface{})
	GetGRPCClient(name string) interface{}
	GetInternalGRPCClient(name string) interface{}
	RegisterNetListener(server BaseGateway)
	GetMQTTClient() mqtt.Client
	GetTGBot() *tgbotapi.BotAPI
	MQTTPublish(topic string, obj interface{})
}

type ForGateway interface {
	GetContext() context.Context
	GetConfig() *config.Config
	GetModel(name mongodb.CollectionName) mongodb.Model
	GetTaskManager(name string) client.Client
}

func (f *frmwrk) OnStartup(fn func(ctx context.Context)) {
	fn(f.ctx)
}

func (f *frmwrk) OnFinish(fn func(ctx context.Context)) {
	fn(f.ctx)
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

func (f *frmwrk) RegisterGRPCServer(server BaseTwirpServer, path string, middlewares ...func(handler http.Handler) http.Handler) {
	if f.grpcRoutes == nil {
		f.grpcRoutes = http.NewServeMux()
	}
	hnd := server.(http.Handler)
	for _, middleware := range middlewares {
		hnd = middleware(hnd)
	}
	f.grpcRoutes.Handle(server.PathPrefix()+path, hnd)
}

func (f *frmwrk) RegisterInternalGRPCServer(server BaseTwirpServer, path string, middlewares ...func(handler http.Handler) http.Handler) {
	if f.internalGrpcRoutes == nil {
		f.internalGrpcRoutes = http.NewServeMux()
	}
	hnd := server.(http.Handler)
	for _, middleware := range middlewares {
		hnd = middleware(hnd)
	}
	f.internalGrpcRoutes.Handle(server.PathPrefix()+path, hnd)
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

func (f *frmwrk) RegisterWorker(namespace string, queue string, options worker.Options, fn func(wrk worker.Worker)) {
	cl := f.GetTaskManager(namespace)
	wrk := worker.New(cl, queue, options)
	fn(wrk)
	f.Workers.Store(namespace+"-"+queue, wrk)
}

func (f *frmwrk) GetContext() context.Context {
	return f.ctx
}

func (f *frmwrk) GetMQTTClient() mqtt.Client {
	return f.MqttClient
}

func (f *frmwrk) MQTTPublish(topic string, obj interface{}) {
	if f.MqttClient != nil {
		f.MqttClient.Publish(topic, 0, false, mqtt2.ConvertForMQTT(obj))
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
		log.Fatal().Err(fmt.Errorf("can`t load grpc client: %s", name))
		return nil
	}
	return grpcClient
}

func (f *frmwrk) GetInternalGRPCClient(name string) interface{} {
	internalGRPCClient, ok := f.InternalGRPCClients.Load(name)
	if !ok {
		log.Fatal().Err(fmt.Errorf("can`t load grpc client: %s", name))
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

func (f *frmwrk) GetTaskManager(name string) client.Client {
	model, ok := f.TaskManagers.Load(name)
	if !ok {
		return nil
	}
	return model.(client.Client)
}

func (f *frmwrk) GetTGBot() *tgbotapi.BotAPI {
	return f.TGBot
}

// Run main livecycle
func (f *frmwrk) Run() error {
	wg, _ := errgroup.WithContext(context.Background())

	defer func() {
		if f.MqttClient != nil && f.MqttClient.IsConnected() {
			defer f.MqttClient.Disconnect(0)
		}
		if f.onShutdown != nil {
			f.onShutdown(f.ctx)
		}
	}()

	if f.config.Redis.URI != "" {
		f.PubSub = redisPS.New(f.config)
	}

	if f.config.Mongo.URI != "" {
		cl, err := mongodb.NewClient(f.ctx, f.config.Mongo.URI)
		if err != nil {
			return err
		}
		defer cl.Disconnect(f.ctx)
		for _, entity := range f.config.Mongo.Entities {
			f.Models.Store(entity.Collection.String(), mongodb.New(cl, f.config.Mongo.DatabaseName, entity.Collection))
			if entity.WithMemstore {
				f.Stores[entity.Collection] = &sync.Map{}
			}
		}
	}

	if f.onStartup != nil {
		f.onStartup(f.ctx)
	}

	if f.Gateway != nil {
		f.Gateway.SetFramework(f)
		wg.Go(func() error {
			if f.config.Gateway.Addr == "" {
				fmt.Println("Stack started")
			} else {
				fmt.Printf("Gateway started at %s\n", f.config.Gateway.Addr)
			}
			f.Gateway.Start(f.ctx)
			return nil
		})
	}
	for _, entity := range f.config.Mongo.Entities {
		if !entity.WithMemstore && !entity.WithReplication {
			continue
		}
		mdl, ok := f.Models.Load(entity.Collection.String())
		if !ok {
			return errors.New("model is not set")
		}

		cursor, err := mdl.(mongodb.Model).GetCollection().Find(f.ctx, bson.M{})
		if err != nil {
			return err
		}
		var results []bson.M
		if err = cursor.All(context.TODO(), &results); err != nil {
			log.Fatal().Err(err)
		}
		for _, result := range results {
			data, _ := bson.Marshal(result)
			bus.Ring(listenerReplicationStarted, mongodb.ReplicationEvent{
				Type:           mongodb.ReplicationEventInsert,
				CollectionName: entity.Collection,
				Data:           data,
			})
		}

		wg.Go(func() error {
			cs, err := mdl.(mongodb.Model).GetCollection().Watch(context.TODO(), entity.ReplicationQuery)
			if err != nil {
				panic(err)
			}
			defer cs.Close(context.TODO())
			for cs.Next(context.TODO()) {
				event, err := mongodb.WatchEventHandler(cs)
				if err != nil {
					return err
				}

				if err := bus.Ring(listenerReplication, mongodb.ReplicationEvent{
					Type:           event.OperationType,
					CollectionName: entity.Collection,
					Data:           event.Data,
				}); err != nil {
					return err
				}

				log.Printf("Event: %s, %s, %s", event.ID, event.Collection, event.Data)
			}
			if err := cs.Err(); err != nil {
				return err
			}
			return nil
		})
	}

	if f.grpcRoutes != nil {
		wg.Go(func() error {
			corsWrapper := cors.New(cors.Options{
				AllowedOrigins: []string{"*"},
				AllowedMethods: []string{"OPTIONS", "POST"},
				AllowedHeaders: []string{"Content-Type", "Authorization"},
			})
			fmt.Printf("GRPC server started at %s\n", f.config.Server.Addr)
			server := http.Server{Addr: f.config.Server.Addr, Handler: corsWrapper.Handler(f.grpcRoutes)}
			lnr, err := net.Listen("tcp4", server.Addr)
			if err != nil {
				return err
			}
			return server.Serve(lnr)
		})
	}

	if f.internalGrpcRoutes != nil {
		wg.Go(func() error {
			fmt.Printf("INTERNAL GRPC server started at %s\n", f.config.Server.InternalAddr)
			server := http.Server{Addr: f.config.Server.InternalAddr, Handler: f.internalGrpcRoutes}
			lnr, err := net.Listen("tcp4", server.Addr)
			if err != nil {
				return err
			}
			return server.Serve(lnr)
		})
	}

	f.Workers.Range(func(key, value any) bool {
		wg.Go(func() error {
			err := value.(worker.Worker).Run(worker.InterruptCh())
			if err != nil {
				return err
			}
			return nil
		})
		return true
	})

	return wg.Wait()
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

	if frm.config.MQTTClient.URI != "" {
		mqtOpt := mqtt.NewClientOptions()
		mqtOpt.AddBroker(fmt.Sprintf("tcp://%s", frm.config.MQTTClient.URI))
		//mqtOpt.SetDefaultPublishHandler(messagePubHandler)
		mqtOpt.SetUsername(frm.config.MQTTClient.Username)
		mqtOpt.SetPassword(frm.config.MQTTClient.Password)
		mqtOpt.SetClientID(frm.config.MQTTClient.ClientId)
		mqClient := mqtt.NewClient(mqtOpt)
		if token := mqClient.Connect(); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}

		frm.MqttClient = mqClient
	}

	if frm.config.Temporal.URI != "" {
		for _, ns := range frm.config.Temporal.Namespaces {
			tm, err := tasks.New(frm.config.Temporal.URI, ns)
			if err != nil {
				log.Fatal().Err(err)
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
