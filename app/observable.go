package app

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	"github.com/segmentio/kafka-go"
	"reactive-kafka-boilerplate/app/cfg"
	"reactive-kafka-boilerplate/app/events"
	"reactive-kafka-boilerplate/app/rxerrs"
)

func (app *Instance) SendErrorEvent(eventName string, value []byte) error {
	ctx := context.Background()
	kafkaProducer := events.GetKafkaProducer(cfg.GetConfig())

	return kafkaProducer.WriteMessages(ctx, kafka.Message{
		Value: value,
		Topic: eventName,
	})
}

func recoverHandling(app *Instance, ctx *events.Ctx) {
	if p := recover(); p != nil {
		rxerrs.HandleError(p.(error), ctx, app)
	}
}

func Observe(observable rxgo.Observable, app *Instance) {
	<-observable.
		Map(events.AcquireKafkaCtx, rxgo.WithCPUPool()).
		ForEach(
			func(item interface{}) {
				ctx := item.(*events.Ctx)
				defer recoverHandling(app, ctx)

				handler := app.handlers[ctx.EventName]
				handler.HandlerFunction(ctx, handler.ContainerFabric)
			},
			func(err error) {
				rxerrs.HandleError(err, nil, app)
			},
			func() {})
}
