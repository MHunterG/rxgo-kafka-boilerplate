package handlers

import (
	"context"
	"github.com/MHunterG/rxgo-kafka-boilerplate/app"
	"github.com/MHunterG/rxgo-kafka-boilerplate/app/cfg"
	"github.com/MHunterG/rxgo-kafka-boilerplate/app/events"
	"github.com/segmentio/kafka-go"
)

type EchoContainer struct {
	KafkaWriter *kafka.Writer
}

func (container EchoContainer) Destroy() {}

func InitEchoContainer() app.Container {
	return EchoContainer{
		KafkaWriter: events.GetKafkaProducer(cfg.GetConfig()),
	}
}

func Echo(c *events.Ctx, fabric app.ContainerFabric) {
	container := fabric().(EchoContainer)
	defer container.Destroy()

	err := container.KafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Key:   c.OriginalMessage.Key,
		Value: c.OriginalMessage.Value,
		Topic: c.EventName + "-response",
	})

	if err != nil {
		panic(err)
	}
}

func RegisterEcho(a *app.Instance) {
	a.RegisterHandler("boilerplate-echo", Echo, InitEchoContainer)
}
