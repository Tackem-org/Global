package masterData

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"sync"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/system/setupData"
)

type Infostruct struct {
	URL             string
	Port            uint32
	RegistrationKey string
}

var (
	setupOnce sync.Once
	UP        helpers.Locker = helpers.Locker{Label: "Master"}
	Info      Infostruct
)

const (
	defaultURL  string = "127.0.0.1"
	defaultPort uint32 = 50000
)

//TODO Split this down for better testing somehow then test
func Setup() {
	setupOnce.Do(func() {
		UP.Down()
		if file, err := ioutil.ReadFile(setupData.Data.MasterConf); err == nil {
			_ = json.Unmarshal([]byte(file), &Info)
			return
		}

		if val, present := os.LookupEnv("URL"); present {
			Info.URL = val
		} else {
			Info.URL = defaultURL
		}
		if val, present := os.LookupEnv("PORT"); present {
			p, errp := strconv.Atoi(val)
			if errp != nil {
				Info.Port = defaultPort
			} else {
				Info.Port = uint32(p)
			}
		} else {
			Info.Port = defaultPort
		}
		if val, present := os.LookupEnv("REGKEY"); present {
			Info.RegistrationKey = val
		}

		file, _ := json.MarshalIndent(Info, "", " ")
		_ = ioutil.WriteFile(setupData.Data.MasterConf, file, 0644)

	})

}
