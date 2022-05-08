package masterData

import (
	"encoding/json"
	"os"
	"strconv"
	"sync"

	"github.com/Tackem-org/Global/helpers"
)

var (
	setupOnce      sync.Once
	UP             helpers.Locker = helpers.Locker{Label: "Master"}
	Info           Infostruct
	ConnectionInfo ConnectionInfostruct
	Setup          = setup
)

func ResetFuncs() {
	Setup = setup
}

const (
	DefaultURL  string = "127.0.0.1"
	DefaultPort uint32 = 50000
)

func setup(masterConf string) bool {
	setupOnce.Do(func() {
		UP.Down()
	})

	Info.URL = DefaultURL
	Info.Port = DefaultPort
	if grabFromFile(masterConf) {
		return true
	}
	return grabFromEnv()

}

func grabFromFile(masterConf string) bool {
	if masterConf == "" {
		return false
	}

	if file, err := os.ReadFile(masterConf); err == nil {
		return json.Unmarshal([]byte(file), &Info) == nil
	}
	return false
}

func grabFromEnv() bool {
	urlVal, urlPresent := os.LookupEnv("URL")
	portVal, portPresent := os.LookupEnv("PORT")
	keyVal, keyPresent := os.LookupEnv("REGKEY")
	os.Unsetenv("REGKEY")
	os.Unsetenv("URL")
	os.Unsetenv("PORT")

	if !keyPresent || !urlPresent || !portPresent {
		return false
	}
	Info.RegistrationKey = keyVal
	Info.URL = urlVal
	p, errp := strconv.Atoi(portVal)
	if errp == nil {
		Info.Port = uint32(p)
	}

	return true
}
