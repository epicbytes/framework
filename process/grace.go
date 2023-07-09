package process

import (
	"context"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Operation is a clean up function on shutting down
type Operation func(ctx context.Context) error

// gracefulShutdown waits for termination syscalls and doing clean up operations after received it
func GracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]Operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Info().Msg("shutting down")

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Info().Msgf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Info().Msgf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Info().Msgf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Info().Msgf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
