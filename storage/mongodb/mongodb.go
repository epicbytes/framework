package mongodb

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CollectionName string

func (c CollectionName) String() string {
	return string(c)
}

type ReplicationEventType string

const (
	ReplicationEventInsert ReplicationEventType = "insert"
	ReplicationEventUpdate ReplicationEventType = "update"
	ReplicationEventDelete ReplicationEventType = "delete"
)

var DefaultReplicationPipeline = mongo.Pipeline{
	bson.D{
		{"$match", bson.D{
			{"operationType", bson.D{
				{"$in", bson.A{"insert", "update", "delete"}},
			}}}}},
}

func (r ReplicationEventType) String() string {
	return string(r)
}

type ReplicationEvent struct {
	Id             string
	Type           ReplicationEventType
	CollectionName CollectionName
	Data           []byte
}

func (r *ReplicationEvent) Decode(object interface{}) error {
	return bson.Unmarshal(r.Data, object)
}

type ModelEntity struct {
	Collection       CollectionName
	WithMemstore     bool
	WithReplication  bool
	ReplicationQuery mongo.Pipeline
}

type model struct {
	client         *mongo.Client
	collection     *mongo.Collection
	databaseName   string
	collectionName CollectionName
	counterModel   CounterModel
}

func (m *model) initModel(client *mongo.Client, dbName string, collectionName CollectionName) {
	m.client = client
	m.collectionName = collectionName
	m.collection = client.Database(dbName).Collection(collectionName.String())
	m.counterModel = CounterModel{collection: client.Database(dbName).Collection(collectionName.String())}
}

func (m *model) GetCollection() *mongo.Collection {
	return m.collection
}

func (m *model) GetNextID(ctx context.Context) (id uint32, err error) {
	return m.counterModel.NextID(ctx, m.collectionName.String())
}

type Model interface {
	GetCollection() *mongo.Collection
}

func New(client *mongo.Client, dbName string, collectionName CollectionName) Model {
	mdl := &model{}
	mdl.initModel(client, dbName, collectionName)
	return mdl
}

type WatchEvent struct {
	Collection    string
	ID            string
	OperationType ReplicationEventType
	Data          []byte
}

func WatchEventHandler(cs *mongo.ChangeStream) (*WatchEvent, error) {
	var event bson.M
	if err := cs.Decode(&event); err != nil {
		return nil, err
	}

	fullDocument, err := json.Marshal(event["fullDocument"])
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	documentKey := event["documentKey"].(primitive.M)
	ns := event["ns"].(primitive.M)

	return &WatchEvent{
		Collection:    ns["coll"].(string),
		ID:            documentKey["_id"].(primitive.ObjectID).Hex(),
		OperationType: ReplicationEventType(event["operationType"].(string)),
		Data:          fullDocument,
	}, nil
}
