package requiredServices

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

var (
	rs []*RequiredService
)

type RequiredService struct {
	ServiceName string
	ServiceType string
	BaseID      string
	Key         string
	IPAddress   string
	Port        uint32
	SingleRun   bool
}

func Get(serviceName string, serviceType string) *RequiredService {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.requiredServices.Get")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] serviceName=%s, serviceType=%s", serviceName, serviceType)
	for _, s := range rs {
		if s.ServiceName == serviceName && s.ServiceType == serviceType {
			return s
		}
	}
	return nil
}

func GetByBaseID(baseID string) *RequiredService {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.requiredServices.GetByBaseID")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
	for _, s := range rs {
		if s.BaseID == baseID {
			return s
		}
	}
	return nil
}

func Add(r *RequiredService) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.requiredServices.Add")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] RequiredService=%+v", r)
	for _, s := range rs {
		if s.BaseID == r.BaseID {
			return
		}
	}
	rs = append(rs, r)
}

func Remove(baseID string) bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.requiredServices.Remove")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
	for index, s := range rs {
		if s.BaseID == baseID {
			rs = append(rs[:index], rs[index+1:]...)
			return true
		}
	}
	return false
}
