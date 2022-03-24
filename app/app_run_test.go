package app_test

import (
	"github.com/stretchr/testify/assert"
	"reactive-kafka-boilerplate/app"
	"reactive-kafka-boilerplate/app/cfg"
	"reactive-kafka-boilerplate/handlers"
	"testing"
	"time"
)

func TestInstance_Run(t *testing.T) {
	c := cfg.GetConfig()
	a := app.New(c)
	handlers.RegisterEcho(a)

	go a.Run()
	defer func() {
		a.Shutdown()
		time.Sleep(1 * time.Second)
		assert.False(t, a.IsRunning())
	}()

	time.Sleep(1 * time.Second)
	assert.True(t, a.IsRunning())
}
