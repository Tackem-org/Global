package helpers_test

import (
	"os"
	"strings"
	"testing"

	"github.com/Tackem-org/Global/helpers"
	"github.com/stretchr/testify/assert"
)

func TestHasEnv(t *testing.T) {
	helpers.ResetDockerFunctions()
	_, err := os.Stat("/.dockerenv")
	env := !os.IsNotExist(err)
	assert.Equal(t, env, helpers.HasEnv())
}

func TestHasCGroup(t *testing.T) {
	helpers.ResetDockerFunctions()
	f, _ := os.ReadFile("/proc/self/cgroup")
	cGroup := strings.Contains(string(f), "docker")
	assert.Equal(t, cGroup, helpers.HasCGroup())
}

func TestInDocker(t *testing.T) {
	helpers.ResetDockerFunctions()
	_, err := os.Stat("/.dockerenv")
	env := !os.IsNotExist(err)
	f, _ := os.ReadFile("/proc/self/cgroup")
	cGroup := strings.Contains(string(f), "docker")

	assert.Equal(t, env || cGroup, helpers.InDocker())
}

func TestHasSocket(t *testing.T) {
	helpers.ResetDockerFunctions()
	_, err := os.Stat("/var/run/docker.sock")
	hasSocket := !os.IsNotExist(err)
	assert.Equal(t, hasSocket, helpers.HasSocket())
}
