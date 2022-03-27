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
)

const (
	DefaultURL  string = "127.0.0.1"
	DefaultPort uint32 = 50000
)

func Setup(masterConf string) bool {
	setupOnce.Do(func() {
		UP.Down()
	})

	Info.URL = DefaultURL
	Info.Port = DefaultPort
	if grabFromFile(masterConf) {
		return true
	}

	if grabFromEnv() {
		return saveToFile(masterConf)
	}
	return false
}

func grabFromFile(masterConf string) bool {
	if masterConf == "" {
		return false
	}

	if file, err := os.ReadFile(masterConf); err == nil {
		err = json.Unmarshal([]byte(file), &Info)
		if err == nil {
			return true
		}
	}
	return false
}

func grabFromEnv() bool {
	urlVal, urlPresent := os.LookupEnv("URL")
	portVal, portPresent := os.LookupEnv("PORT")
	keyVal, keyPresent := os.LookupEnv("REGKEY")

	if !keyPresent {
		return false
	}
	Info.RegistrationKey = keyVal

	if urlPresent {
		Info.URL = urlVal
	}
	if portPresent {
		p, errp := strconv.Atoi(portVal)
		if errp == nil {
			Info.Port = uint32(p)
		}
	}

	os.Unsetenv("REGKEY")
	os.Unsetenv("URL")
	os.Unsetenv("PORT")
	return true
}

func saveToFile(masterConf string) bool {
	if masterConf == "" {
		return false
	}
	file, _ := json.MarshalIndent(Info, "", " ")
	return os.WriteFile(masterConf, file, 0644) == nil
}
