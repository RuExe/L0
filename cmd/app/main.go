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
	//orders := []domain.Order{fakefactory.CreateFakeOrder()}
	//
	//mapa := map[string]*domain.Order{
	//	"test": &orders[0],
	//}
	//
	//mapa["test"].TrackNumber = "asdadafasfas"
	//
	//fmt.Println(orders)

	config := config.GetConfig()

	store := pstgr.NewStore(config)
	if err := store.Open(); err != nil {
		log.Fatal(err)
	}

	var rep repository.OrderStorage = store.Order()
	//var storage repository.OrderStorage = repository.NewCacheOrderStorage(&rep)

	subscriber := subscriber.NewSubscriber(config, &rep)
	go subscriber.Subscribe()

	var orderService = service.NewOrderService(&rep)

	server := server.NewServer(config, orderService)
	go server.Start()

	runtime.Goexit()
}
