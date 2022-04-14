package regClient

import (
	"context"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	pb "github.com/Tackem-org/Global/pb/regclient"
	"github.com/Tackem-org/Global/system/channels"
	"github.com/Tackem-org/Global/system/masterData"
)

func (r *RegClientServer) MasterGoingDown(ctx context.Context, in *pb.MasterGoingDownRequest) (*pb.MasterGoingDownResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.MasterGoingDownResponse{Success: false, ErrorMessage: err}, nil
	}

	switch in.GetReason() {
	case pb.MasterGoingDownReason_FullShutdown:
		logging.Info("Master Is Down, Shutting down this service")
		channels.Root.Shutdown <- true

	case pb.MasterGoingDownReason_Shutdown:
		logging.Info("Master Is Down")
		if masterData.UP.Check() {
			masterData.UP.Down()
		} else {
			return &pb.MasterGoingDownResponse{
				Success:      false,
				ErrorMessage: "master already down",
			}, nil
		}
	default:
		return &pb.MasterGoingDownResponse{
			Success:      false,
			ErrorMessage: "unknown reason for going down",
		}, nil
	}

	return &pb.MasterGoingDownResponse{
		Success: true,
	}, nil
}
