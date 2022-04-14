package regClient

import (
	"context"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	pb "github.com/Tackem-org/Global/pb/regclient"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
)

func (r *RegClientServer) MasterBackUp(ctx context.Context, in *pb.MasterBackUpRequest) (*pb.MasterBackUpResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.MasterBackUpResponse{Success: false, ErrorMessage: err}, nil
	}

	if masterData.UP.Check() {
		return &pb.MasterBackUpResponse{
			Success:      false,
			ErrorMessage: "master already marked as up",
			Active:       setupData.Active,
			RegData:      setupData.Data.RegisterProto(),
		}, nil
	}

	logging.Info("Master Is Back Up")
	masterData.UP.Up()

	return &pb.MasterBackUpResponse{
		Success: true,
		Active:  setupData.Active,
		RegData: setupData.Data.RegisterProto(),
	}, nil

}
