package app

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"reactive-kafka-boilerplate/app/cfg"
	"reactive-kafka-boilerplate/app/events"
	"syscall"
)

type Instance struct {
	config      *cfg.Config
	kafkaReader *kafka.Reader
	handlers    map[string]handler
	isRunning   bool
}

func (app *Instance) IsRunning() bool {
	return app.isRunning
}

func New(config *cfg.Config) *Instance {
	return &Instance{
		config:    config,
		handlers:  map[string]handler{},
		isRunning: false,
	}
}

func (app *Instance) Run() {
	app.isRunning = true
	app.kafkaReader = events.GetKafkaConsumer(app.config, app.GetTopics())
	observable := rxgo.Defer([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		for {
			msg, err := app.kafkaReader.ReadMessage(context.Background())
			if err != nil {
				log.Warning(err)
				break
			}

			next <- rxgo.Item{V: &msg}
			fmt.Println(msg.Offset)
		}

		fmt.Println("Stopping...")

		if err := app.kafkaReader.Close(); err != nil {
			log.Fatal("failed to close reader:", err)
		}

		app.isRunning = false
	}})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		app.Shutdown()
	}()

	Observe(observable, app)
}

func (app *Instance) Shutdown() {
	err := app.kafkaReader.Close()
	if err != nil {
		log.Warning(err)
		return
	}
}
