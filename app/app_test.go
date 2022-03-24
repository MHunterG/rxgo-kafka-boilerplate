package app

import (
	"github.com/stretchr/testify/assert"
	"reactive-kafka-boilerplate/app/cfg"
	"testing"
)

func TestNew(t *testing.T) {
	c := cfg.GetConfig()
	a := New(c)
	assert.Equal(t, a.config, c)
}
