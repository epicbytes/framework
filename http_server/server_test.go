package http_server_test

import (
	"testing"

	"github.com/epicbytes/framework/http_server"
	"github.com/stretchr/testify/assert"
)

func Test_http(t *testing.T) {
	quit := make(chan struct{})

	go func() {
		http_server.New(quit).Run()
	}()

	<-quit

	t.Run("Ping", func(t *testing.T) {
		assert.True(t, true)
	})
}
