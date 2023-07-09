package service

import (
	"context"
	"github.com/epicbytes/framework/config"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"net"
	"net/http"
	"sync"
	"time"
)

type GRPCService struct {
	ctx             context.Context
	mux             sync.RWMutex
	requestTime     time.Time
	Config          *config.Config
	GrpcMultiplexer *http.ServeMux
	server          http.Server
	listener        net.Listener
}

func (t *GRPCService) Init(ctx context.Context) error {
	t.ctx = ctx
	log.Debug().Msg("INITIAL GRPC SERVICE")
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{
			"Accept-Encoding",
			"Content-Encoding",
			"Content-Type",
			"Connect-Protocol-Version",
			"Connect-Timeout-Ms",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Grpc-Timeout",
			"X-Grpc-Web",
			"X-User-Agent",
		},
		ExposedHeaders: []string{
			"Content-Encoding",
			"Connect-Content-Encoding",
			"Grpc-Status",
			"Grpc-Message",
		},
	})
	log.Info().Msgf("GRPC server started at %s", t.Config.Server.Addr)
	t.server = http.Server{Addr: t.Config.Server.Addr, Handler: corsWrapper.Handler(t.GrpcMultiplexer)}
	lnr, err := net.Listen("tcp4", t.server.Addr)
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

func (t *GRPCService) Ping(context.Context) error {
	//log.Debug().Msg("PING GRPC")
	return nil
}

func (t *GRPCService) Close() error {
	log.Debug().Msg("CLOSE GRPC")
	err := t.server.Shutdown(t.ctx)
	if err != nil {
		return err
	}
	return nil
}
