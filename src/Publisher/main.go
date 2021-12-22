package main

import (
	"flag"
	"github.com/nats-io/nats.go"
	"log"
	"os"
)

func usage() {
	log.Printf("Usage: nats-pub <subject> <msg>\n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func main() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		showUsageAndExit(1)
	}

	nc, err := nats.Connect(*urls)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subj, msg := args[0], []byte(args[1])
	nc.Publish(subj, msg)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
}
