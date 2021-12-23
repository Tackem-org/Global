package system

import (
	"context"
	"sync"

	pb "github.com/Tackem-org/Proto/pb/checker"
)

type CheckerServer struct {
	mu sync.RWMutex
	pb.UnimplementedCheckerServer
}

func NewCheckerServer() *CheckerServer {
	return &CheckerServer{}
}

func (c *CheckerServer) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	return &pb.HealthCheckResponse{Healthy: true}, nil
}
