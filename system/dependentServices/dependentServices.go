package dependentServices

import (
	"fmt"
	"sync"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

var (
	mu sync.RWMutex
	ds []*DependentService
)

type DependentService struct {
	UP          helpers.Locker
	setupOnce   sync.Once
	ServiceName string
	ServiceType string
	ServiceID   uint64
	BaseID      string
	Key         string
	IPAddress   string
	Port        uint32
	SingleRun   bool
}

func (ds *DependentService) setup() {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.dependentServices.DependentService{%s %s}.setup", ds.ServiceType, ds.ServiceName)
	ds.UP.Up()
	ds.UP.Label = fmt.Sprintf("[Dependent] %s %s", ds.ServiceType, ds.ServiceName)
}

func GetActive() []*DependentService {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.dependentServices.GetActive")
	mu.RLock()
	defer mu.RUnlock()
	var rd []*DependentService

	for _, s := range ds {
		if s.UP.Check() {
			continue
		}
		rd = append(rd, s)
	}
	return rd
}

func GetByBaseID(baseID string) *DependentService {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.dependentServices.GetByBaseID")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
	mu.RLock()
	defer mu.RUnlock()
	for _, s := range ds {
		if s.BaseID == baseID {
			return s
		}
	}
	return nil
}

func Add(d *DependentService) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.dependentServices.Add")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] dependentService=%+v", d)
	mu.Lock()
	defer mu.Unlock()
	for _, s := range ds {
		if s.BaseID == d.BaseID {
			return
		}
	}
	d.setupOnce.Do(d.setup)
	ds = append(ds, d)
}

func Remove(baseID string) bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.dependentServices.Remove")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
	mu.Lock()
	defer mu.Unlock()
	for index, s := range ds {
		if s.BaseID == baseID {
			ds = append(ds[:index], ds[index+1:]...)
			return true
		}
	}
	return false
}

func Up(baseID string) bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.dependentServices.Up")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
	mu.Lock()
	defer mu.Unlock()
	for _, s := range ds {
		if s.BaseID == baseID {
			s.UP.Up()
			return true
		}
	}
	return false
}

func Down(baseID string) bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.dependentServices.Down")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
	mu.Lock()
	defer mu.Unlock()
	for _, s := range ds {
		if s.BaseID == baseID {
			s.UP.Down()
			return true
		}
	}
	return false
}
