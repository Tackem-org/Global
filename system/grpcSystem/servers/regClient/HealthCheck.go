package regClient

import (
	"context"

	"github.com/Tackem-org/Global/healthCheck"
	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	pb "github.com/Tackem-org/Global/pb/regclient"
	"github.com/Tackem-org/Global/system/masterData"
)

func (r *RegClientServer) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.HealthCheckResponse{Success: false, ErrorMessage: err}, nil
	}
	logging.Info("Health Check OK")
	return &pb.HealthCheckResponse{
		Success: true,
		Healthy: healthCheck.Healthy(),
		Issues:  healthCheck.Issues(),
	}, nil
}
