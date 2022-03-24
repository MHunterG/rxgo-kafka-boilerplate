package app

import "reactive-kafka-boilerplate/app/events"

type Container interface {
	Destroy()
}

type ContainerFabric func() Container
type HandlerFunction func(ctx *events.Ctx, fabric ContainerFabric)

type handler struct {
	ContainerFabric ContainerFabric
	HandlerFunction HandlerFunction
}

func (app *Instance) RegisterHandler(eventName string, function HandlerFunction, fabric ContainerFabric) {
	app.handlers[eventName] = handler{
		ContainerFabric: fabric,
		HandlerFunction: function,
	}
}

func (app *Instance) GetTopics() []string {
	topics := make([]string, 0, len(app.handlers))
	for k := range app.handlers {
		topics = append(topics, k)
	}
	return topics
}
