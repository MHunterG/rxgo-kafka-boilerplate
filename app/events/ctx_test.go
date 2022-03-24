package events_test

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"reactive-kafka-boilerplate/app/events"
	"testing"
)

func TestAcquireKafkaCtx(t *testing.T) {
	kafkaMessage := &kafka.Message{
		Topic: "test-topic",
		Value: []byte("test-value"),
	}

	ctxRaw, err := events.AcquireKafkaCtx(context.TODO(), kafkaMessage)
	if err != nil {
		panic(err)
	}

	ctx := ctxRaw.(*events.Ctx)

	assert.Equal(t, ctx.OriginalMessage, kafkaMessage)
	assert.Equal(t, ctx.Payload, kafkaMessage.Value)
	assert.Equal(t, ctx.EventName, kafkaMessage.Topic)
}

func TestCtx_PayloadParser(t *testing.T) {
	ctx := events.Ctx{
		Payload:   []byte(`{"test": "value"}`),
		EventName: "",
	}

	type testPayload struct {
		Test string `json:"test"`
	}

	p := testPayload{}

	err := ctx.PayloadParser(&p)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, p.Test, "value")
}
