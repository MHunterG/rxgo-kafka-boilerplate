package cfg_test

import (
	"github.com/stretchr/testify/assert"
	cfg3 "reactive-kafka-boilerplate/app/cfg"
	"testing"
)

func TestGetConfig(t *testing.T) {
	cfg := cfg3.GetConfig()
	cfg2 := cfg3.GetConfig()

	assert.Equal(t, cfg, cfg2)
}
