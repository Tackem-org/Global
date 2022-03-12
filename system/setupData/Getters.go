package setupData

import (
	"fmt"
)

func (d *SetupData) GetPath(path string) *PathItem {
	d.mu.RLock()
	defer d.mu.RUnlock()
	for _, p := range d.Paths {
		if p.Path == path {
			return p
		}
	}
	return nil
}

func (d *SetupData) GetAdminPath(path string) *AdminPathItem {
	d.mu.RLock()
	defer d.mu.RUnlock()
	for _, ap := range d.AdminPaths {
		if ap.Path == path {
			return ap
		}
	}
	return nil
}

func (d *SetupData) GetSocket(command string) *SocketItem {
	d.mu.RLock()
	defer d.mu.RUnlock()
	for _, s := range d.Sockets {
		if s.Command == command {
			return s
		}
	}
	return nil
}

func (d *SetupData) Name() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if d.ServiceType == "system" {
		return d.ServiceName
	}
	return fmt.Sprintf("%s %s", d.ServiceType, d.ServiceName)
}
