package helpers

import (
	"sync"
	"time"

	lock "github.com/viney-shih/go-lock"
)

type Locker struct {
	l         lock.Mutex
	setupOnce sync.Once
	Label     string
}

func (m *Locker) setup() {
	m.l = lock.NewCASMutex()
}

func (m *Locker) Down() {
	m.setupOnce.Do(m.setup)
	m.l.Lock()
}

func (m *Locker) Up() {
	m.setupOnce.Do(m.setup)
	m.l.Unlock()
}

func (m *Locker) Check() bool {
	m.setupOnce.Do(m.setup)
	if m.l.TryLock() {
		m.l.Unlock()
		return true
	}
	return false
}

func (m *Locker) Wait(timeout time.Duration) bool {
	m.setupOnce.Do(m.setup)
	if m.l.TryLockWithTimeout(timeout) {
		m.l.Unlock()
		return true
	}
	return false
}
