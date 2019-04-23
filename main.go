package main

import (
	"context"

	"github.com/mheers/emitter/broker"
	emitterConfig "github.com/mheers/emitter/config"
)

func startEmitterSvc() {
	cfg := &emitterConfig.Config{
		ListenAddr: ":8081",
		License:    "J9VGSRu7nCDkU8VPhRPYG2cJm4n6V-9bAAAAAAAAAAI",
	}

	// Start new service
	svc, err := broker.NewService(context.Background(), cfg)
	if err != nil {
		panic(err.Error())
	}

	// Listen and serve
	svc.Listen()
}

func main() {
	startEmitterSvc()
}
