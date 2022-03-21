package regClient

import (
	"context"

	"github.com/Tackem-org/Global/helpers"
	pb "github.com/Tackem-org/Global/pb/regclient"
	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/masterData"
)

func (r *RegClientServer) RemoveDependent(ctx context.Context, in *pb.RemoveDependentRequest) (*pb.RemoveDependentResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.RemoveDependentResponse{Success: false, ErrorMessage: err}, nil
	}

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
