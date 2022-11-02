package main

import (
	"context"
	tradeCrypto "github.com/samuelmachado/go-trade-crypto"
	"math/rand"
	"time"

	//tradeCrypto "github.com/samuelmachado/go-trade-crypto"
	"github.com/samuelmachado/go-trade-crypto/internal/container"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	rand.Seed(time.Now().UnixNano())

	if os.Getenv("ENV") == "" || os.Getenv("ENV") == "dev" {
		tradeCrypto.LoadEnvFromFile("env/application.env")
		log.Println("Environment variables have not been set on the OS. We load from a file, this should only be used for local development")
	}

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
