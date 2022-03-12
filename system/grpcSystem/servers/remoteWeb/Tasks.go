package remoteWeb

import (
	"context"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
)

func (r *RemoteWebServer) Tasks(ctx context.Context, in *pb.TasksRequest) (*pb.TasksResponse, error) {
	logging.Info("Master Has Requested Tasks")
	t := setupData.Data.TaskGrabber()
	if len(t) == 0 {
		logging.Info("No Tasks To Send (sending empty list)")
	} else {
		logging.Info("Sending Master Tasks")

	}
	return &pb.TasksResponse{
		Tasks: t,
	}, nil

}
