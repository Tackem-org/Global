package remoteWeb

import (
	"context"

	"github.com/Tackem-org/Global/helpers"
	pb "github.com/Tackem-org/Global/pb/remoteweb"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
)

func (r *RemoteWebServer) Notifications(ctx context.Context, in *pb.NotificationsRequest) (*pb.NotificationsResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.NotificationsResponse{Success: false, ErrorMessage: err}, nil
	}
	return &pb.NotificationsResponse{
		Success:       true,
		Notifications: setupData.Data.NotificationGrabber(),
	}, nil

}
