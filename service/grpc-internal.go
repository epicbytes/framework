package service

import (
	"context"
	"github.com/epicbytes/framework/config"
	"github.com/rs/zerolog/log"
	"net"
	"net/http"
	"sync"
	"time"
)

type GRPCInternalService struct {
	ctx                     context.Context
	mux                     sync.RWMutex
	requestTime             time.Time
	Config                  *config.Config
	GrpcInternalMultiplexer *http.ServeMux
	server                  http.Server
	listener                net.Listener
}

func (t *GRPCInternalService) Init(ctx context.Context) error {
	t.ctx = ctx
	log.Debug().Msg("INITIAL GRPC Internal SERVICE")
	log.Info().Msgf("GRPC internal server started at %s", t.Config.Server.InternalAddr)
	server := http.Server{Addr: t.Config.Server.InternalAddr, Handler: t.GrpcInternalMultiplexer}
	lnr, err := net.Listen("tcp4", server.Addr)
	if err != nil {
		return err
	}
	go func() {
		err = t.server.Serve(lnr)
		if err != nil {
			log.Error().Err(err)
			return
		}
	}()
	return nil
}

func (t *GRPCInternalService) Ping(context.Context) error {
	return nil
}

func (t *GRPCInternalService) Close() error {
	log.Debug().Msg("CLOSE GRPC Internal")
	err := t.server.Shutdown(t.ctx)
	if err != nil {
		return err
	}
	return nil
}
