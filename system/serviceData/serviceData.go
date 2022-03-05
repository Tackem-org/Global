package masterData

import (
	"github.com/Tackem-org/Global/helpers"
)

type ServiceData struct {
	UP          helpers.Locker
	ServiceType string
	ServiceName string
	ServiceID   uint64
	BaseID      string
	URL         string
	Port        uint32
	Key         string
	SingleRun   bool
}
