package dependentServices

import (
	"sync"
)

var (
	mu sync.RWMutex
	ds []*DependentService
)

func GetActive() []*DependentService {
	mu.RLock()
	defer mu.RUnlock()
	var rd []*DependentService

	for _, s := range ds {
		if !s.UP.Check() {
			continue
		}
		rd = append(rd, s)
	}
	return rd
}

func GetByBaseID(baseID string) *DependentService {
	mu.RLock()
	defer mu.RUnlock()
	for _, s := range ds {
		if s.BaseID == baseID {
			return s
		}
	}
	return nil
}

func Add(d *DependentService) bool {
	mu.Lock()
	defer mu.Unlock()
	for _, s := range ds {
		if s.BaseID == d.BaseID {
			return false
		}
	}
	d.setupOnce.Do(d.setup)
	ds = append(ds, d)
	return true
}

func Remove(baseID string) bool {
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
	mu.Lock()
	defer mu.Unlock()
	for _, s := range ds {
		if s.BaseID == baseID {
			if !s.UP.Check() {
				s.UP.Up()
				return true
			}
			return false
		}
	}
	return false
}

func Down(baseID string) bool {
	mu.Lock()
	defer mu.Unlock()
	for _, s := range ds {
		if s.BaseID == baseID {
			if s.UP.Check() {

				s.UP.Down()
				return true
			}
			return false
		}
	}
	return false
}
