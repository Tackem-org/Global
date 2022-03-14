package requiredServices

import (
	"fmt"
	"sync"

	"github.com/Tackem-org/Global/helpers"
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
	rs.UP.Label = fmt.Sprintf("[Required] %s %s", rs.ServiceType, rs.ServiceName)
}

func Get(serviceName string, serviceType string) *RequiredService {
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
	mu.RLock()
	defer mu.RUnlock()
	for _, s := range rs {
		if s.BaseID == baseID {
			return s
		}
	}
	return nil
}

func Add(r *RequiredService) bool {
	mu.Lock()
	defer mu.Unlock()
	for _, s := range rs {
		if s.BaseID == r.BaseID {
			return false
		}
	}
	r.setupOnce.Do(r.setup)
	rs = append(rs, r)
	return true
}

func Remove(baseID string) bool {
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
	mu.Lock()
	defer mu.Unlock()
	for _, s := range rs {
		if s.BaseID == baseID {
			if !s.UP.Check() {
				s.UP.Up()
				return true
			}
		}
	}
	return false
}

func Down(baseID string) bool {
	mu.Lock()
	defer mu.Unlock()
	for _, s := range rs {
		if s.BaseID == baseID {
			if s.UP.Check() {
				s.UP.Down()
				return true
			}
		}
	}
	return false
}
