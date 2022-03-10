package requiredServices

import (
	"fmt"
	"sync"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

var (
	mu sync.RWMutex
	rs []*RequiredService
)

type RequiredService struct {
	UP          helpers.Locker
	setupOnce   sync.Once
	ServiceName string
	ServiceType string
	ServiceID   uint64
	BaseID      string
	Key         string
	IPAddress   string
	Port        uint32
}

func (rs *RequiredService) setup() {
	rs.UP.Up()
	rs.UP.Label = fmt.Sprintf("[Required] %s %s", rs.ServiceType, rs.ServiceName)
}

func Get(serviceName string, serviceType string) *RequiredService {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.requiredServices.Get")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] serviceName=%s, serviceType=%s", serviceName, serviceType)
	mu.RLock()
	defer mu.RUnlock()
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
	mu.RLock()
	defer mu.RUnlock()
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
	mu.Lock()
	defer mu.Unlock()
	for _, s := range rs {
		if s.BaseID == r.BaseID {
			return
		}
	}
	r.setupOnce.Do(r.setup)
	rs = append(rs, r)
}

func Remove(baseID string) bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.requiredServices.Remove")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
	mu.Lock()
	defer mu.Unlock()
	for index, s := range rs {
		if s.BaseID == baseID {
			rs = append(rs[:index], rs[index+1:]...)
			return true
		}
	}
	return false
}

func Up(baseID string) bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.requiredServices.Up")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
	mu.Lock()
	defer mu.Unlock()
	for _, s := range rs {
		if s.BaseID == baseID {
			s.UP.Up()
			return true
		}
	}
	return false
}

func Down(baseID string) bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.requiredServices.Down")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] baseID=%s", baseID)
	mu.Lock()
	defer mu.Unlock()
	for _, s := range rs {
		if s.BaseID == baseID {
			s.UP.Down()
			return true
		}
	}
	return false
}
