package regClient

import (
	"context"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) MasterBackUp(ctx context.Context, in *pb.MasterBackUpRequest) (*pb.MasterBackUpResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.servers.regClient.RegClientServer{}.MasterBackUp")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] ctx in=%+v", in)

	ok, err := checkKey(ctx)
	if ok {
		logging.Info("Master Is Back Up")
		masterData.UP.Up()
	}
	return &pb.MasterBackUpResponse{
		Success:      ok,
		ErrorMessage: err,
		Active:       setupData.Active,
	}, nil
}
