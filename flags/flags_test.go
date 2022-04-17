package flags_test

import (
	"testing"

	"github.com/Tackem-org/Global/flags"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestConfigFolder(t *testing.T) {
	folder := "/config/"
	pflag.Set("config", folder)
	assert.Equal(t, folder, flags.ConfigFolder())
}

func TestLogFile(t *testing.T) {
	folder := "/logs/"
	pflag.Set("log", folder)
	assert.Equal(t, folder, flags.LogFolder())
}

func TestVerbose(t *testing.T) {
	assert.False(t, flags.Verbose())
	pflag.Set("verbose", "true")
	assert.True(t, flags.Verbose())
}

func TestParse(t *testing.T) {
	assert.NotPanics(t, func() { flags.Parse() })
}
