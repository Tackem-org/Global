package regClient

import (
	"context"

	"github.com/Tackem-org/Global/helpers"
	pb "github.com/Tackem-org/Global/pb/regclient"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
)

func (r *RegClientServer) RemoveRequired(ctx context.Context, in *pb.RemoveRequiredRequest) (*pb.RemoveRequiredResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.RemoveRequiredResponse{Success: false, ErrorMessage: err}, nil
	}
	if requiredServices.Remove(in.BaseId) {
		return &pb.RemoveRequiredResponse{
			Success: true,
		}, nil
	}
	return &pb.RemoveRequiredResponse{
		Success:      false,
		ErrorMessage: "service not found",
	}, nil
}
