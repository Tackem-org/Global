package regClient

import (
	"context"

	"github.com/Tackem-org/Global/system/dependentServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) DependentDown(ctx context.Context, in *pb.DependentDownRequest) (*pb.DependentDownResponse, error) {
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
