package regClient

import (
	"context"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/dependentServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) RemoveDependant(ctx context.Context, in *pb.RemoveDependantRequest) (*pb.RemoveDependantResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.servers.regClient.RegClientServer{}.RemoveDependant")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] ctx in=%+v", in)
	if dependentServices.Remove(in.BaseId) {
		return &pb.RemoveDependantResponse{
			Success: true,
		}, nil
	}
	return &pb.RemoveDependantResponse{
		Success:      false,
		ErrorMessage: "Service Not Found",
	}, nil
}
