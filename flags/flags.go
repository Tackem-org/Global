package flags

import (
	flag "github.com/spf13/pflag"
)

var (
	configFolder = flag.StringP("config", "c", "/config/", "Config Location")
	logFolder    = flag.StringP("log", "l", "/logs/", "Log Location")
	verbose      = flag.BoolP("verbose", "v", false, "Outputs the log to the screen")
)

func ConfigFolder() string {
	return *configFolder
}

func LogFolder() string {
	return *logFolder
}

func Verbose() bool {
	return *verbose
}
