package flags_test

import (
	"testing"

	"github.com/Tackem-org/Global/flags"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestRegistrationFile(t *testing.T) {
	folder := "/config/test.json"
	pflag.Set("regfile", folder)
	assert.Equal(t, folder, flags.RegistrationFile())
}
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

func TestVersion(t *testing.T) {
	assert.False(t, flags.Version())
	pflag.Set("version", "true")
	assert.True(t, flags.Version())
}

func TestParse(t *testing.T) {
	assert.NotPanics(t, func() { flags.Parse() })
}
