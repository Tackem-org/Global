package regClient

import (
	"context"

	"github.com/Tackem-org/Global/healthCheck"
	"github.com/Tackem-org/Global/logging"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	logging.Info("Health Check OK")
	return &pb.HealthCheckResponse{
		Healthy: healthCheck.Healthy(),
		Issues:  healthCheck.Issues(),
	}, nil
}
