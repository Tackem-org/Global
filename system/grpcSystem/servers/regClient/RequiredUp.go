package regClient

import (
	"context"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
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
		ErrorMessage: "required Service Not Found",
	}, nil
}
