package cfg_test

import (
	cfg3 "github.com/MHunterG/rxgo-kafka-boilerplate/app/cfg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfig(t *testing.T) {
	cfg := cfg3.GetConfig()
	cfg2 := cfg3.GetConfig()

	assert.Equal(t, cfg, cfg2)
}
