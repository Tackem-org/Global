package regClient

import (
	"context"

	"github.com/Tackem-org/Global/system/requiredServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) RemoveRequired(ctx context.Context, in *pb.RemoveRequiredRequest) (*pb.RemoveRequiredResponse, error) {
	if requiredServices.Remove(in.BaseId) {
		return &pb.RemoveRequiredResponse{
			Success: true,
		}, nil
	}
	return &pb.RemoveRequiredResponse{
		Success:      false,
		ErrorMessage: "Service Not Found",
	}, nil
}
