package regClient

import (
	"context"

	"github.com/Tackem-org/Global/channels"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/masterData"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) MasterGoingDown(ctx context.Context, in *pb.MasterGoingDownRequest) (*pb.MasterGoingDownResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.servers.regClient.RegClientServer{}.MasterGoingDown")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] ctx in=%+v", in)

	ok, err := checkKey(ctx)
	if ok {
		switch in.GetReason() {
		case pb.MasterGoingDownReason_FullShutdown:
			logging.Info("Master Is Down, Shutting down this service")
			channels.Root.Shutdown <- true

		case pb.MasterGoingDownReason_Shutdown:
			logging.Info("Master Is Down")
			masterData.UP.Down()
		}
	}
	return &pb.MasterGoingDownResponse{
		Success:      ok,
		ErrorMessage: err,
	}, nil
}
