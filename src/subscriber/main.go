package main

import (
	"L0/core"
	"L0/subscriber/classes"
	"L0/subscriber/repositories"
	"L0/subscriber/services"
	"log"
	"runtime"
)

func main() {
	config := core.GetConfig()

	store := repositories.NewStore(config)
	if err := store.Open(); err != nil {
		log.Fatal(err)
	}

	var rep repositories.OrderStorage = store.Order()
	var storage repositories.OrderStorage = repositories.NewCacheOrderStorage(&rep)

	subscriber := classes.NewSubscriber(config, &storage)
	go subscriber.Subscribe()

	var orderService = services.NewOrderService(&storage)

	server := classes.NewServer(config, orderService)
	go server.Start()

	runtime.Goexit()
}
