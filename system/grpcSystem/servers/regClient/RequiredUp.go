package regClient

import (
	"context"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/requiredServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) RequiredUp(ctx context.Context, in *pb.RequiredUpRequest) (*pb.RequiredUpResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.servers.regClient.RegClientServer{}.RequiredUp")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] ctx in=%+v", in)
	if requiredServices.Up(in.BaseId) {
		return &pb.RequiredUpResponse{
			Success: true,
		}, nil
	}
	return &pb.RequiredUpResponse{
		Success:      false,
		ErrorMessage: "required Service Not Found",
	}, nil
}
