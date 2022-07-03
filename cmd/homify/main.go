package main

import (
	"os"
	"os/signal"

	"go.uber.org/zap"

	"github.com/MovieStoreGuy/homify/internal/context"
	"github.com/MovieStoreGuy/homify/internal/worker"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)
	defer done()

	ctx, done = context.WithCancel(ctx)
	defer done()

	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	w := worker.NewPool[any]()
	w.Add(worker.JobFunc[any](func(ctx context.Context) (any, error) {

		return nil, nil
	}))

	if err := w.Do(ctx); err != nil {
		log.Panic("Issue with service", zap.Error(err))
	}
	log.Info("Application shutdown")
}
