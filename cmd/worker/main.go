package main

import (
	"context"
	"github.com/samuelmachado/go-trade-crypto/internal/container"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	ctx, dep, err := container.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	run(ctx, dep)
}

func run(ctx context.Context, dep *container.Dependency) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer signal.Stop(interrupt)

	dep.Components.Log.Info(
		ctx,
		"Starting application",
	)

	<-interrupt
	dep.Components.Log.Info(ctx, "Stopping application")

}
