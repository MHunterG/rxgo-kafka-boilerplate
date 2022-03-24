package app_test

import (
	"github.com/MHunterG/rxgo-kafka-boilerplate/app"
	"github.com/MHunterG/rxgo-kafka-boilerplate/app/cfg"
	"github.com/MHunterG/rxgo-kafka-boilerplate/handlers"
	"github.com/stretchr/testify/assert"
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
