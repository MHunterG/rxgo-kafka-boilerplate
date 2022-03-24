package cfg

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	KafkaGroupID string
	KafkaHost    string `env:"KAFKA_HOST,required"`
}

var configInstance *Config = nil
var once sync.Once

func NewConfig() *Config {
	cfg := Config{
		KafkaGroupID: "boilerplate",
	}
	return &cfg
}

func GetConfig() *Config {
	once.Do(func() {
		configInstance = NewConfig()
		if err := env.Parse(configInstance); err != nil {
			log.Error(err)
			panic(err)
		}
	})

	return configInstance
}
