package regClient

import (
	"context"

	"github.com/Tackem-org/Global/helpers"
	pb "github.com/Tackem-org/Global/pb/regclient"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
)

func (r *RegClientServer) RequiredUp(ctx context.Context, in *pb.RequiredUpRequest) (*pb.RequiredUpResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.RequiredUpResponse{Success: false, ErrorMessage: err}, nil
	}
	if requiredServices.Up(in.BaseId) {
		return &pb.RequiredUpResponse{
			Success: true,
		}, nil
	}
	return &pb.RequiredUpResponse{
		Success:      false,
		ErrorMessage: "required service not found",
	}, nil
}
