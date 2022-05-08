package flags

import (
	"github.com/spf13/pflag"
	flag "github.com/spf13/pflag"
)

var (
	registrationFile = pflag.StringP("regfile", "r", "data.json", "registration file")
	configFolder     = flag.StringP("config", "c", "/config/", "config location")
	logFolder        = flag.StringP("log", "l", "/logs/", "log location")
	version          = flag.BoolP("version", "v", false, "outputs the current version")
)

func RegistrationFile() string {
	return *registrationFile
}

func ConfigFolder() string {
	return *configFolder
}

func LogFolder() string {
	return *logFolder
}

func Version() bool {
	return *version
}

func Parse() {
	pflag.Parse()

}
