package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"sharelock/config"
	"sharelock/pkg/locker"
	"sharelock/server"
)

func main() {
	// global context
	globalCtx, cancelGlobalCtx := context.WithCancel(context.Background())
	cfg := config.ReadConfig()

	// locker
	locker := locker.NewLocker()
	go locker.Start(globalCtx)

	serversList := make([]server.Server, 0)
	if cfg.GrpcServer != nil && cfg.GrpcServer.Enable {
		serversList = append(serversList, server.NewGrpcServer(globalCtx, cfg.GrpcServer, locker))
	}
	if cfg.HttpServer != nil && cfg.HttpServer.Enable {
		serversList = append(serversList, server.NewHttpServer(globalCtx, cfg.HttpServer, locker))
	}

	// start all the servers
	for i := range serversList {
		go serversList[i].Start()
	}

	// Wait for Control C to exit
	ch := make(chan os.Signal, 3)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch

	// Finally, we stop the server
	log.Println("[WARN] stopping ShareLock")
	for i := range serversList {
		serversList[i].Stop()
	}
	cancelGlobalCtx()
}
