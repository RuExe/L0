package main

import (
	"L0/internal/config"
	"L0/internal/fakefactory"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	config := config.GetConfig()

	nc, err := nats.Connect(config.SubscriberConfig.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subj := config.SubscriberConfig.Subject
	msg, _ := json.Marshal(fakefactory.CreateFakeOrder())

	nc.Publish(config.SubscriberConfig.Subject, msg)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
}
