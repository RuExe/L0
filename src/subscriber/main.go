package main

import (
	"L0/core"
	"L0/subscriber/classes"
	"runtime"
)

func main() {
	config := core.NewConfig()
	server := classes.NewServer(config)

	go server.Start()
	//go classes.Subscribe(config, service)

	runtime.Goexit()
}
