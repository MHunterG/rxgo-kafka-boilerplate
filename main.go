package main

import (
	"github.com/MHunterG/rxgo-kafka-boilerplate/app"
	"github.com/MHunterG/rxgo-kafka-boilerplate/app/cfg"
	"github.com/MHunterG/rxgo-kafka-boilerplate/handlers"
	"github.com/MHunterG/rxgo-kafka-boilerplate/logs"
)

func main() {
	logs.Setup()
	config := cfg.GetConfig()
	a := app.New(config)
	handlers.RegisterEcho(a)

	a.Run()
}
