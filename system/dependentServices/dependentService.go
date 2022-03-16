package dependentServices

import (
	"fmt"
	"sync"

	"github.com/Tackem-org/Global/helpers"
)

type DependentService struct {
	UP          helpers.Locker
	setupOnce   sync.Once
	ServiceName string
	ServiceType string
	ServiceID   uint64
	BaseID      string
	Key         string
	URL         string
	Port        uint32
	SingleRun   bool
}

func (ds *DependentService) setup() {
	ds.UP.Label = fmt.Sprintf("[Dependent] %s %s", ds.ServiceType, ds.ServiceName)
}

func (ds *DependentService) Address() string {
	if ds.Port == 0 {
		return ds.URL
	}
	return fmt.Sprintf("%s:%d", ds.URL, ds.Port)
}
