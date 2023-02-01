package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigInit(t *testing.T) {
	envInit()
	assert.Equal(t, "localhost", EnvConfig.PGSQL_HOST)
}
