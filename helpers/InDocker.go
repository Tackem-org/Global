package helpers

import (
	"io/ioutil"
	"os"
	"strings"
)

var (
	HasEnv    = hasEnv
	HasCGroup = hasCGroup
	InDocker  = inDocker
	HasSocket = hasSocket
)

func ResetDockerFunctions() {
	HasEnv = hasEnv
	HasCGroup = hasCGroup
	InDocker = inDocker
	HasSocket = hasSocket
}

func hasEnv() bool {
	_, err := os.Stat("/.dockerenv")
	return !os.IsNotExist(err)
}

func hasCGroup() bool {
	f, _ := ioutil.ReadFile("/proc/self/cgroup")
	return strings.Contains(string(f), "docker")
}

func inDocker() bool {
	return HasEnv() || HasCGroup()
}

func hasSocket() bool {
	_, err := os.Stat("/var/run/docker.sock")
	return !os.IsNotExist(err)
}
