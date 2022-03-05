package regClient

import (
	"context"

	"github.com/Tackem-org/Global/healthCheck"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

func (r *RegClientServer) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.servers.regClient.RegClientServer{}.HealthCheck")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] ctx in=%+v", in)

	logging.Info("Health Check OK")
	return &pb.HealthCheckResponse{
		Healthy: healthCheck.Healthy(),
		Issues:  healthCheck.Issues(),
	}, nil
}
