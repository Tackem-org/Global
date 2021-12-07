package indocker

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func hasEnv() bool {
	_, err := os.Stat("/.dockerenv")
	return !os.IsNotExist(err)
}

func hasCGroup() bool {
	f, _ := ioutil.ReadFile("/proc/self/cgroup")
	return strings.Contains(string(f), "docker")
}

func Check() bool {
	if hasEnv() || hasCGroup() {
		return true
	}
	fmt.Println("program not in a docker container please run in a container (see readme.md for more information")
	return false
}
