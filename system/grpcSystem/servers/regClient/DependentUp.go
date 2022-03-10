package regClient

import (
	"context"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/dependentServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) DependentUp(ctx context.Context, in *pb.DependentUpRequest) (*pb.DependentUpResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.servers.regClient.RegClientServer{}.DependentUp")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] ctx in=%+v", in)
	if dependentServices.Up(in.BaseId) {
		return &pb.DependentUpResponse{
			Success: true,
		}, nil
	}
	return &pb.DependentUpResponse{
		Success:      false,
		ErrorMessage: "Dependent Service Not Found",
	}, nil

}
