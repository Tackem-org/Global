package regClient

import (
	"context"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	pb "github.com/Tackem-org/Global/pb/regclient"
	"github.com/Tackem-org/Global/system/channels"
	"github.com/Tackem-org/Global/system/masterData"
)

func (r *RegClientServer) ShutdownService(ctx context.Context, in *pb.ShutdownServiceRequest) (*pb.ShutdownServiceResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.ShutdownServiceResponse{Success: false, ErrorMessage: err}, nil
	}

	logging.Info("Master Told Me To Shutdown")
	channels.Root.Shutdown <- true

	return &pb.ShutdownServiceResponse{
		Success: true,
	}, nil
}
