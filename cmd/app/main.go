package main

import (
	"L0/cmd/server"
	"L0/cmd/subscriber"
	"L0/internal/config"
	"L0/internal/repository"
	"L0/internal/repository/pstgr"
	"L0/internal/service"
	"log"
	"runtime"
)

func main() {
	config := config.GetConfig()

	store := pstgr.NewStore(config)
	if err := store.Open(); err != nil {
		log.Fatal(err)
	}

	var rep repository.OrderStorage = store.Order()
	var storage repository.OrderStorage = repository.NewCacheOrderStorage(&rep)

	subscriber := subscriber.NewSubscriber(config, &storage)
	go subscriber.Subscribe()

	var orderService = service.NewOrderService(&storage)

	server := server.NewServer(config, orderService)
	go server.Start()

	runtime.Goexit()
}
