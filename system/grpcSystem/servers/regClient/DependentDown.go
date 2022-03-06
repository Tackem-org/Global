package regClient

import (
	"context"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/dependentServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) DependentDown(ctx context.Context, in *pb.DependentDownRequest) (*pb.DependentDownResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.servers.regClient.RegClientServer{}.DependentDown")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] ctx in=%+v", in)
	if dependentServices.Down(in.BaseId) {
		return &pb.DependentDownResponse{
			Success: true,
		}, nil
	}
	return &pb.DependentDownResponse{
		Success:      false,
		ErrorMessage: "Dependent Service Not Found",
	}, nil
}
