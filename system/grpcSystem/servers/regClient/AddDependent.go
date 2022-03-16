package regClient

import (
	"context"

	"github.com/Tackem-org/Global/system/dependentServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) AddDependent(ctx context.Context, in *pb.AddDependentRequest) (*pb.AddDependentResponse, error) {
	if s := dependentServices.GetByBaseID(in.BaseId); s != nil {
		return &pb.AddDependentResponse{
			Success: true,
		}, nil
	}

	dependentServices.Add(&dependentServices.DependentService{
		ServiceType: in.Type,
		ServiceName: in.Name,
		ServiceID:   in.Id,
		BaseID:      in.BaseId,
		Key:         in.Key,
		URL:         in.Url,
		Port:        in.Port,
		SingleRun:   in.SingleRun,
	})
	return &pb.AddDependentResponse{
		Success: true,
	}, nil
}
