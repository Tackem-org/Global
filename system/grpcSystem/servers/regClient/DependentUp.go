package regClient

import (
	"context"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/masterData"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) DependentUp(ctx context.Context, in *pb.DependentUpRequest) (*pb.DependentUpResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.DependentUpResponse{Success: false, ErrorMessage: err}, nil
	}

	if dependentServices.Up(in.BaseId) {
		return &pb.DependentUpResponse{
			Success: true,
		}, nil
	}
	return &pb.DependentUpResponse{
		Success:      false,
		ErrorMessage: "Dependent Service Not Found",
	}, nil

}
