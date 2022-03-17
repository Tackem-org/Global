package regClient

import (
	"context"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/masterData"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) AddDependent(ctx context.Context, in *pb.AddDependentRequest) (*pb.AddDependentResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.AddDependentResponse{Success: false, ErrorMessage: err}, nil
	}

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
