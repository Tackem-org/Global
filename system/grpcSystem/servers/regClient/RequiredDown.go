package regClient

import (
	"context"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/requiredServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) RequiredDown(ctx context.Context, in *pb.RequiredDownRequest) (*pb.RequiredDownResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.servers.regClient.RegClientServer{}.RequiredDown")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] ctx in=%+v", in)
	if requiredServices.Down(in.BaseId) {
		return &pb.RequiredDownResponse{
			Success: true,
		}, nil
	}
	return &pb.RequiredDownResponse{
		Success:      false,
		ErrorMessage: "required Service Not Found",
	}, nil
}
