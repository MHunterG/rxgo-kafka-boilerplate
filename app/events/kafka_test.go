package events_test

import (
	"github.com/MHunterG/rxgo-kafka-boilerplate/app/cfg"
	"github.com/MHunterG/rxgo-kafka-boilerplate/app/events"
	"testing"
)

func TestGetKafkaConsumer(t *testing.T) {
	config := cfg.GetConfig()
	consumer := events.GetKafkaConsumer(config, []string{"test-event"})

	err := consumer.Close()
	if err != nil {
		panic(err)
	}
}

func TestGetKafkaProducer(t *testing.T) {
	config := cfg.GetConfig()
	producer := events.GetKafkaProducer(config)

	err := producer.Close()
	if err != nil {
		panic(err)
	}
}
