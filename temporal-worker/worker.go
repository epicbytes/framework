// THIS FILE CREATED WITH GENERATOR DO NOT EDIT!
package temporal_worker

import (
	"github.com/epicbytes/framework/temporal"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
)

type TemporalWorker struct {
	Config   *Config
	Logger   *zap.Logger
	Temporal client.Client
	Worker   worker.Worker
}

func newTemporalWorker(config *Config, logger *zap.Logger, tmprl *temporal.Temporal) *TemporalWorker {
	clientOptions := client.Options{
		Logger: temporal.NewZapAdapter(logger),
	}
	temporalClient, _ := client.Dial(clientOptions)
	wrk := &TemporalWorker{
		Config:   config,
		Logger:   logger,
		Temporal: temporalClient,
		Worker:   worker.New(tmprl.Client, config.TaskQueue, worker.Options{}),
	}
	return wrk
}
func (s *TemporalWorker) StartWorker() error {
	return s.Worker.Run(worker.InterruptCh())
}
func (s *TemporalWorker) StopWorker() error {
	s.Worker.Stop()
	return nil
}
