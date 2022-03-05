package masterData

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"sync"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/setupData"
)

var (
	setupOnce       sync.Once
	UP              helpers.Locker = helpers.Locker{Label: "Master"}
	URL             string
	Port            uint32
	RegistrationKey string
)

const (
	defaultURL  string = "127.0.0.1"
	defaultPort uint32 = 50000
)

func Setup() {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.masterData.Setup")
	setupOnce.Do(func() {
		data := struct {
			URL             string
			Port            uint32
			RegistrationKey string
		}{}
		file, err := ioutil.ReadFile(setupData.Data.MasterConf)
		if err != nil {
			if val, present := os.LookupEnv("URL"); present {
				data.URL = val
			} else {
				data.URL = defaultURL
			}
			if val, present := os.LookupEnv("PORT"); present {
				p, errp := strconv.Atoi(val)
				if errp != nil {
					data.Port = defaultPort
				} else {
					data.Port = uint32(p)
				}
			} else {
				data.Port = defaultPort
			}
			if val, present := os.LookupEnv("REGKEY"); present {
				data.RegistrationKey = val
			}
			file, _ := json.MarshalIndent(data, "", " ")
			_ = ioutil.WriteFile(setupData.Data.MasterConf, file, 0644)
		} else {
			_ = json.Unmarshal([]byte(file), &data)
		}
		Port = data.Port
		URL = data.URL
		RegistrationKey = data.RegistrationKey
		UP.StartDown()
	})

}
