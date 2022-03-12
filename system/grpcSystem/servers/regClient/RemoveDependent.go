package regClient

import (
	"context"

	"github.com/Tackem-org/Global/system/dependentServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) RemoveDependent(ctx context.Context, in *pb.RemoveDependentRequest) (*pb.RemoveDependentResponse, error) {
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
