package tasks

import (
	"go.temporal.io/sdk/client"
)

func New(address string, namespace string) (client.Client, error) {
	c, err := client.Dial(client.Options{
		HostPort:  address,
		Namespace: namespace,
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}
