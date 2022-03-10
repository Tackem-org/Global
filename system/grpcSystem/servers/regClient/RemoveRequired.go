package regClient

import (
	"context"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/requiredServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) RemoveRequired(ctx context.Context, in *pb.RemoveRequiredRequest) (*pb.RemoveRequiredResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.servers.regClient.RegClientServer{}.RemoveRequired")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] ctx in=%+v", in)
	if requiredServices.Remove(in.BaseId) {
		return &pb.RemoveRequiredResponse{
			Success: true,
		}, nil
	}
	return &pb.RemoveRequiredResponse{
		Success:      false,
		ErrorMessage: "Service Not Found",
	}, nil
}
