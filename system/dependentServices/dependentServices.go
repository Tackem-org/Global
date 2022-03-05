package dependentServices

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

type DependentService struct {
	BaseID    string
	Key       string
	IPAddress string
	Port      uint32
	SingleRun bool
}

var (
	ds []*DependentService
)

func GetByBaseID(baseID string) *DependentService {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.dependentServices.GetByBaseID")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
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
	for index, s := range ds {
		if s.BaseID == baseID {
			ds = append(ds[:index], ds[index+1:]...)
			return true
		}
	}
	return false
}
