package events

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

type RawKafkaCtx struct {
	Message *kafka.Message
}

type Ctx struct {
	OriginalMessage *kafka.Message
	Payload         []byte
	EventName       string
}

func (ctx *Ctx) PayloadParser(out interface{}) error {
	err := json.Unmarshal(ctx.Payload, out)
	if err != nil {
		return err
	}

	return nil
}

func AcquireKafkaCtx(_ context.Context, item interface{}) (interface{}, error) {
	message := item.(*kafka.Message)
	c := Ctx{
		Payload:         message.Value,
		EventName:       message.Topic,
		OriginalMessage: message,
	}

	return &c, nil
}
