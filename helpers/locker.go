package helpers

import "sync"

type Locker struct {
	b bool
	l sync.Mutex
}

func (m *Locker) StartDown() {
	m.b = false
	m.l.Lock()
}

func (m *Locker) StartUp() {
	m.b = true
	m.l.Unlock()
}

func (m *Locker) Down() {
	if m.b {
		m.b = false
		m.l.Lock()
	}
}

func (m *Locker) Up() {
	if !m.b {
		m.b = true
		m.l.Unlock()
	}
}

func (m *Locker) Wait() {
	if !m.b {
		m.l.Lock()
		defer m.l.Unlock()
	}
}

func (m *Locker) Check() bool {
	return m.b
}

func (m *Locker) CheckAndWait() bool {
	if !m.b {
		m.l.Lock()
		defer m.l.Unlock()
		return true
	}
	return false
}
