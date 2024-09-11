package bell

import (
	"github.com/epicbytes/framework/auth"
	"github.com/nuttech/bell/v2"
	"go.uber.org/zap"
)

type WatcherMessage struct {
	UserData auth.BaseUserModel
	Data     interface{}
}

type Bell struct {
	Events *bell.Events
	Config *Config
	Logger *zap.Logger
	Done   chan struct{}
}

func newBell(logger *zap.Logger, config *Config) *Bell {
	return &Bell{
		Events: bell.New(),
		Config: config,
		Logger: logger,
		Done:   make(chan struct{}),
	}
}
