package setupData

import (
	"github.com/Tackem-org/Global/helpers"
	pb "github.com/Tackem-org/Proto/pb/registration"
)

func (d *SetupData) RegisterProto() *pb.RegisterRequest {
	return &pb.RegisterRequest{
		ServiceName:       d.ServiceName,
		ServiceType:       d.ServiceType,
		Version:           d.Version.ToProto(),
		Port:              Port,
		Multi:             d.Multi,
		SingleRun:         d.SingleRun,
		StartActive:       d.StartActive,
		ConfigItems:       d.ConfigItems,
		NavItems:          d.NavItems,
		RequiredServices:  d.RequiredServices,
		WebLinkItems:      d.PathsToProtos(),
		AdminWebLinkItems: d.AdminPathsToProtos(),
		WebSocketItems:    d.SocketsToProtos(),
		Groups:            d.Groups,
		Permissions:       d.Permissions,
	}
}
func (d *SetupData) AdminPathsToProtos() []*pb.AdminWebLinkItem {
	d.mu.RLock()
	defer d.mu.RUnlock()
	var r []*pb.AdminWebLinkItem
	for _, p := range d.AdminPaths {
		if helpers.CheckPath(p.Path) {
			r = append(r, &pb.AdminWebLinkItem{
				Path:        p.Path,
				PostAllowed: p.PostAllowed,
				GetDisabled: p.GetDisabled,
			})
		}
	}
	return r
}

func (d *SetupData) PathsToProtos() []*pb.WebLinkItem {
	d.mu.RLock()
	defer d.mu.RUnlock()
	var r []*pb.WebLinkItem
	for _, p := range d.Paths {
		if helpers.CheckPath(p.Path) {
			r = append(r, &pb.WebLinkItem{
				Path:        p.Path,
				Permission:  p.Permission,
				PostAllowed: p.PostAllowed,
				GetDisabled: p.GetDisabled,
			})
		}
	}
	return r
}

func (d *SetupData) SocketsToProtos() []*pb.WebSocketItem {
	d.mu.RLock()
	defer d.mu.RUnlock()
	var r []*pb.WebSocketItem
	for _, p := range d.Sockets {
		r = append(r, &pb.WebSocketItem{
			Command:           p.Command,
			Permission:        p.Permission,
			AdminOnly:         p.AdminOnly,
			RequiredVariables: p.RequiredVariables,
		})

	}
	return r
}
