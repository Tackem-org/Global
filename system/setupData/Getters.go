package setupData

import (
	"fmt"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

func (d *SetupData) GetPath(path string) *PathItem {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.setupData.GetPath")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] path=%s", path)
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
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.setupData.GetAdminPath")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] path=%s", path)
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
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.setupData.GetSocket")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] command=%s", command)
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
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.setupData.Name")
	d.mu.RLock()
	defer d.mu.RUnlock()
	if d.ServiceType == "system" {
		return d.ServiceName
	}
	return fmt.Sprintf("%s %s", d.ServiceType, d.ServiceName)
}
