package setupData

import (
	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	pb "github.com/Tackem-org/Proto/pb/registration"
)

func (s SetupData) AdminPathsToProtos() []*pb.AdminWebLinkItem {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.setupData.AdminPathsToProtos")
	var r []*pb.AdminWebLinkItem
	for _, p := range s.AdminPaths {
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

func (s SetupData) PathsToProtos() []*pb.WebLinkItem {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.setupData.PathsToProtos")
	var r []*pb.WebLinkItem
	for _, p := range s.Paths {
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

func (s SetupData) SocketsToProtos() []*pb.WebSocketItem {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.setupData.SocketsToProtos")
	var r []*pb.WebSocketItem
	for _, p := range s.Sockets {
		//TODO Add In Checks Here
		r = append(r, &pb.WebSocketItem{
			Command:           p.Command,
			Permission:        p.Permission,
			AdminOnly:         p.AdminOnly,
			RequiredVariables: p.RequiredVariables,
		})

	}
	return r
}
