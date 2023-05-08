package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, uri string) (client *mongo.Client, err error) {
	/*cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			fmt.Println(evt.Command)
		},
	}*/
	client, err = mongo.NewClient(options2.Client().ApplyURI(uri) /*SetMonitor(cmdMonitor)*/)
	if err != nil {
		return
	}
	err = client.Connect(ctx)
	if err != nil {
		return
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return
	}
	return
}
