package events

import (
	"github.com/MHunterG/rxgo-kafka-boilerplate/app/cfg"
	"github.com/segmentio/kafka-go"
	"sync"
	"time"
)

var onceGetKafkaProducer sync.Once
var w *kafka.Writer

func GetKafkaProducer(config *cfg.Config) *kafka.Writer {
	onceGetKafkaProducer.Do(func() {
		w = &kafka.Writer{
			Addr:         kafka.TCP(config.KafkaHost),
			BatchTimeout: time.Millisecond * 1,
		}
	})

	return w
}

func GetKafkaConsumer(config *cfg.Config, topics []string) *kafka.Reader {
	p := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{config.KafkaHost},
		GroupID:     config.KafkaGroupID,
		GroupTopics: topics,
		MaxBytes:    int(3e6),
		MaxWait:     10 * time.Millisecond,
	})

	return p
}
