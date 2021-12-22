package main

import (
	"L0/publisher/factories"
	"encoding/json"
	"flag"
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	url := nats.DefaultURL

	log.SetFlags(0)
	flag.Parse()

	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subj := "foo"
	args := flag.Args()

	var msg []byte
	if len(args) == 1 {
		args := flag.Args()
		msg = []byte(args[0])

	} else {
		msg, _ = json.Marshal(factories.CreateFakeOrder())
	}

	nc.Publish(subj, msg)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
}
