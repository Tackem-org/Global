package dependentServices

import (
	"sync"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

var (
	mu sync.RWMutex
	ds []*DependentService
)

//TODO use helper locker to work this magic
type DependentService struct {
	BaseID    string
	Key       string
	IPAddress string
	Port      uint32
	SingleRun bool
	down      bool
}

func (ds DependentService) Active() bool {
	logging.Debugf(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.dependentServices.DependentService{%s}.Active", ds.BaseID)
	return !ds.down
}

func GetActive() []*DependentService {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.dependentServices.GetActive")
	mu.RLock()
	defer mu.RUnlock()
	var rd []*DependentService

	for _, s := range ds {
		if s.down {
			continue
		}
		rd = append(rd, s)
	}
	return rd
}
func GetByBaseID(baseID string) *DependentService {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.dependentServices.GetByBaseID")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
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
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] dependentService=%+v", d)
	mu.Lock()
	defer mu.Unlock()
	for _, s := range ds {
		if s.BaseID == d.BaseID {
			return
		}
	}
	ds = append(ds, d)
}

func Remove(baseID string) bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.dependentServices.Remove")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
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
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
	mu.Lock()
	defer mu.Unlock()
	for _, s := range ds {
		if s.BaseID == baseID {
			s.down = false
			return true
		}
	}
	return false
}

func Down(baseID string) bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.dependentServices.Down")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
	mu.Lock()
	defer mu.Unlock()
	for _, s := range ds {
		if s.BaseID == baseID {
			s.down = true
			return true
		}
	}
	return false
}
