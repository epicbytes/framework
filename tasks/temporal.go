package tasks

import (
	"github.com/epicbytes/framework/logger"
	"go.temporal.io/sdk/client"
)

func New(address string, namespace string) (client.Client, error) {
	logz := logger.New("debug")
	c, err := client.Dial(client.Options{
		HostPort:  address,
		Namespace: namespace,
		Logger:    NewZapAdapter(logz),
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}
