package setupData

import (
	"github.com/Tackem-org/Global/helpers"
	pb "github.com/Tackem-org/Global/pb/registration"
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
		Panels:            d.PanelsToProtos(),
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
				Path:          p.Path,
				PostAllowed:   p.PostAllowed,
				GetDisabled:   p.GetDisabled,
				AllowedPanels: p.AllowedPanels,
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
				Path:          p.Path,
				Permission:    p.Permission,
				PostAllowed:   p.PostAllowed,
				GetDisabled:   p.GetDisabled,
				AllowedPanels: p.AllowedPanels,
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

func (d *SetupData) PanelsToProtos() []*pb.PanelSetup {
	d.mu.RLock()
	defer d.mu.RUnlock()
	var r []*pb.PanelSetup
	for _, p := range d.Panels {
		var tempPL = pb.PanelLayout{
			Width:        p.Layout.Width,
			Height:       p.Layout.Height,
			ScrollWidth:  p.Layout.ScrollWidth,
			ScrollHeight: p.Layout.ScrollHeight,
			TitleBar:     p.Layout.TitleBar,
			Minimise:     p.Layout.Minimise,
			Close:        p.Layout.Close,
		}
		r = append(r, &pb.PanelSetup{
			Name:              p.Name,
			Label:             p.Label,
			Description:       p.Description,
			Layout:            &tempPL,
			AdminOnly:         p.AdminOnly,
			Permission:        p.Permission,
			RequiredVariables: requiredVariablesToProto(p.RequiredVariables),
		})

	}
	return r
}

func requiredVariablesToProto(rvs []RequiredVariable) []*pb.RequiredVariable {
	var r []*pb.RequiredVariable
	for _, v := range rvs {
		r = append(r, requiredVariableToProto(&v))
	}
	return r
}

func requiredVariableToProto(rv *RequiredVariable) *pb.RequiredVariable {
	return &pb.RequiredVariable{
		Name:    rv.Name,
		Options: rv.Options,
	}
}
