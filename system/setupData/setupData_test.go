package setupData_test

import (
	"os"
	"testing"

	"github.com/Tackem-org/Global/system/setupData"
	"github.com/stretchr/testify/assert"
)

func TestFreeTCPPort(t *testing.T) {
	setupData.Port = 50002
	first := setupData.FreeTCPPort()
	assert.NotNil(t, first)
	assert.Equal(t, uint32(50002), setupData.Port)
	second := setupData.FreeTCPPort()
	assert.NotNil(t, second)
	assert.Equal(t, uint32(50003), setupData.Port)
	os.Setenv("BIND", "localhost")
	defer os.Unsetenv("ENV_VAR")
	third := setupData.FreeTCPPort()
	assert.NotNil(t, third)
	first.Close()
	second.Close()
	third.Close()
}
