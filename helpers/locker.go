package helpers

import (
	"sync"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

type Locker struct {
	setupOnce sync.Once
	Label     string
	l         chan struct{}
}

func (m *Locker) setup() {
	m.l = make(chan struct{}, 1)
}

func (m *Locker) Down() { //lock
	logging.Debugf(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.Down", m.Label)
	m.setupOnce.Do(m.setup)
	m.l <- struct{}{}
}

func (m *Locker) Up() { //unlock
	logging.Debugf(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.Up", m.Label)
	m.setupOnce.Do(m.setup)
	<-m.l

}

func (m *Locker) Check() bool {
	logging.Debugf(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.Check", m.Label)
	m.setupOnce.Do(m.setup)
	select {
	case m.l <- struct{}{}:
		<-m.l
		return true
	default:
		return false
	}
}

func (m *Locker) Wait(timeout time.Duration) bool {
	logging.Debugf(debug.FUNCTIONCALLS|debug.HELPERLOCKER, "[FUNCTIONCALL] Global.helpers.Locker{%s}.TimeoutCheck", m.Label)
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] timeout=%d", timeout)
	m.setupOnce.Do(m.setup)
	select {
	case m.l <- struct{}{}:
		<-m.l
		return true
	case <-time.After(timeout):
		return false
	}
}
