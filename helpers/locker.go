package helpers

import (
	"sync"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
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
	logging.Debug(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.Down", m.Label)
	m.setupOnce.Do(m.setup)
	m.l.Lock()
}

func (m *Locker) Up() {
	logging.Debug(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.Up", m.Label)
	m.setupOnce.Do(m.setup)
	m.l.Unlock()
}

func (m *Locker) Check() bool {
	logging.Debug(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.Check", m.Label)
	m.setupOnce.Do(m.setup)
	if m.l.TryLock() {
		m.l.Unlock()
		return true
	}
	return false
}

func (m *Locker) Wait(timeout time.Duration) bool {
	logging.Debug(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.TimeoutCheck", m.Label)
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] timeout=%d", timeout)
	m.setupOnce.Do(m.setup)
	if m.l.TryLockWithTimeout(timeout) {
		m.l.Unlock()
		return true
	}
	return false
}
