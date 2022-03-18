package regClient

import (
	"context"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) RequiredDown(ctx context.Context, in *pb.RequiredDownRequest) (*pb.RequiredDownResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.RequiredDownResponse{Success: false, ErrorMessage: err}, nil
	}
	if requiredServices.Down(in.BaseId) {
		return &pb.RequiredDownResponse{
			Success: true,
		}, nil
	}
	return &pb.RequiredDownResponse{
		Success:      false,
		ErrorMessage: "Required Service Not Found",
	}, nil
}
