package masterData

import (
	"fmt"
)

type Infostruct struct {
	URL             string
	Port            uint32
	RegistrationKey string
}

func (is *Infostruct) Address() string {
	if is.Port == 0 {
		return is.URL
	}
	return fmt.Sprintf("%s:%d", is.URL, is.Port)
}

type ConnectionInfostruct struct {
	Key string
	IP  string
}

func (ci *ConnectionInfostruct) CheckKey(key string) bool {
	return key == ci.Key
}

func (ci *ConnectionInfostruct) CheckIP(ipAddress string) bool {
	return ipAddress == ci.IP
}
