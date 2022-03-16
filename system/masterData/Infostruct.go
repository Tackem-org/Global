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
