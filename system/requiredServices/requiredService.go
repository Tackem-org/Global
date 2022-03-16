package requiredServices

import (
	"fmt"
	"sync"

	"github.com/Tackem-org/Global/helpers"
)

type RequiredService struct {
	UP          helpers.Locker
	setupOnce   sync.Once
	ServiceName string
	ServiceType string
	ServiceID   uint64
	BaseID      string
	Key         string
	URL         string
	Port        uint32
}

func (rs *RequiredService) setup() {
	rs.UP.Label = fmt.Sprintf("[Required] %s %s", rs.ServiceType, rs.ServiceName)
}

func (rs *RequiredService) Address() string {
	if rs.Port == 0 {
		return rs.URL
	}
	return fmt.Sprintf("%s:%d", rs.URL, rs.Port)
}
