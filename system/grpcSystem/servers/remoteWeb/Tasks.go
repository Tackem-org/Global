package remoteWeb

import (
	"context"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
)

func (r *RemoteWebServer) Tasks(ctx context.Context, in *pb.TasksRequest) (*pb.TasksResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.TasksResponse{Success: false, ErrorMessage: err}, nil
	}
	return &pb.TasksResponse{
		Success: true,
		Tasks:   setupData.Data.TaskGrabber(),
	}, nil

}
