package regClient

import (
	"context"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/dependentServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) RemoveDependent(ctx context.Context, in *pb.RemoveDependentRequest) (*pb.RemoveDependentResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.servers.regClient.RegClientServer{}.RemoveDependent")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] ctx in=%+v", in)
	if dependentServices.Remove(in.BaseId) {
		return &pb.RemoveDependentResponse{
			Success: true,
		}, nil
	}
	return &pb.RemoveDependentResponse{
		Success:      false,
		ErrorMessage: "Service Not Found",
	}, nil
}
