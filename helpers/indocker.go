package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

func hasEnv() bool {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[helpers.hasEnv() bool]")
	_, err := os.Stat("/.dockerenv")
	return !os.IsNotExist(err)
}

func hasCGroup() bool {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[helpers.hasCGroup() bool]")
	f, _ := ioutil.ReadFile("/proc/self/cgroup")
	return strings.Contains(string(f), "docker")
}

func InDockerCheck() bool {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[helpers.InDockerCheck() bool]")
	if hasEnv() || hasCGroup() {
		return true
	}
	fmt.Println("program not in a docker container please run in a container (see readme.md for more information")
	return false
}
