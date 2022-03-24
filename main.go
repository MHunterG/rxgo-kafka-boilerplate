package main

import (
	"reactive-kafka-boilerplate/app"
	"reactive-kafka-boilerplate/app/cfg"
	"reactive-kafka-boilerplate/handlers"
	"reactive-kafka-boilerplate/logs"
)

func main() {
	logs.Setup()
	config := cfg.GetConfig()
	a := app.New(config)
	handlers.RegisterEcho(a)

	a.Run()
}
