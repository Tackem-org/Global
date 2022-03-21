package regClient

import (
	"context"

	"github.com/Tackem-org/Global/helpers"
	pb "github.com/Tackem-org/Global/pb/regclient"
	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/masterData"
)

func (r *RegClientServer) DependentDown(ctx context.Context, in *pb.DependentDownRequest) (*pb.DependentDownResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.DependentDownResponse{Success: false, ErrorMessage: err}, nil
	}

	if dependentServices.Down(in.BaseId) {
		return &pb.DependentDownResponse{
			Success: true,
		}, nil
	}
	return &pb.DependentDownResponse{
		Success:      false,
		ErrorMessage: "Dependent Service Not Found",
	}, nil
}
