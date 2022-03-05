package regClient

import (
	"context"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/dependentServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) AddDependant(ctx context.Context, in *pb.AddDependantRequest) (*pb.AddDependantResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.servers.regClient.RegClientServer{}.AddDependant")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] ctx in=%+v", in)
	if s := dependentServices.GetByBaseID(in.BaseId); s != nil {
		return &pb.AddDependantResponse{
			Success: true,
		}, nil
	}

	dependentServices.Add(&dependentServices.DependentService{
		BaseID:    in.BaseId,
		Key:       in.Key,
		IPAddress: in.IpAddress,
		Port:      in.Port,
		SingleRun: in.SingleRun,
	})
	return &pb.AddDependantResponse{
		Success: true,
	}, nil
}
