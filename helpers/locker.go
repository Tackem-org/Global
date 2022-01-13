package helpers

import (
	"sync"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

type Locker struct {
	b bool
	l sync.Mutex
}

func (m *Locker) StartDown() {
	logging.Debug(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "CALLED:[helpers.(m *Locker) StartDown()]")
	m.b = false
	m.l.Lock()
}

func (m *Locker) StartUp() {
	logging.Debug(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "CALLED:[helpers.(m *Locker) StartUp()]")
	m.b = true
	m.l.Unlock()
}

func (m *Locker) Down() {
	logging.Debug(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "CALLED:[helpers.(m *Locker) Down()]")
	if m.b {
		m.b = false
		m.l.Lock()
	}
}

func (m *Locker) Up() {
	logging.Debug(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "CALLED:[helpers.(m *Locker) Up()]")
	if !m.b {
		m.b = true
		m.l.Unlock()
	}
}

func (m *Locker) Wait() {
	logging.Debug(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "CALLED:[helpers.(m *Locker) Wait()]")
	if !m.b {
		m.l.Lock()
		defer m.l.Unlock()
	}
}

func (m *Locker) Check() bool {
	logging.Debug(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "CALLED:[helpers.(m *Locker) Check() bool]")
	return m.b
}

func (m *Locker) CheckAndWait() bool {
	logging.Debug(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "CALLED:[helpers.(m *Locker) CheckAndWait() bool]")
	if !m.b {
		m.l.Lock()
		defer m.l.Unlock()
		return true
	}
	return false
}
