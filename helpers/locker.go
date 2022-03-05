package helpers

import (
	"sync"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

type Locker struct {
	Label string
	b     bool
	l     sync.Mutex
}

func (m *Locker) StartDown() {
	logging.Debugf(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.StartDown", m.Label)
	m.b = false
	m.l.Lock()
}

func (m *Locker) StartUp() {
	logging.Debugf(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.StartUp", m.Label)
	m.b = true
	m.l.Unlock()
}

func (m *Locker) Down() {
	logging.Debugf(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.Down", m.Label)
	if m.b {
		m.b = false
		m.l.Lock()
	}
}

func (m *Locker) Up() {
	logging.Debugf(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.Up", m.Label)
	if !m.b {
		m.b = true
		m.l.Unlock()
	}
}

func (m *Locker) Wait() {
	logging.Debugf(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.Wait", m.Label)
	if !m.b {
		m.l.Lock()
		defer m.l.Unlock()
	}
}

func (m *Locker) Check() bool {
	logging.Debugf(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.Check", m.Label)
	return m.b
}

func (m *Locker) CheckAndWait() bool {
	logging.Debugf(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.CheckAndWait", m.Label)
	if !m.b {
		m.l.Lock()
		defer m.l.Unlock()
		return true
	}
	return false
}
