package app

import (
	"github.com/MHunterG/rxgo-kafka-boilerplate/app/cfg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	c := cfg.GetConfig()
	a := New(c)
	assert.Equal(t, a.config, c)
}
