package system

import (
	"context"
	"sync"

	pb "github.com/Tackem-org/Proto/pb/regclient"
)

type RegClientServer struct {
	mu sync.RWMutex
	pb.UnimplementedRegClientServer
}

func NewRegClientServer() *RegClientServer {
	return &RegClientServer{}
}

func (r *RegClientServer) MasterGoingDown(ctx context.Context, in *pb.GoingDownRequest) (*pb.Empty, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	MUp.Down()
	return &pb.Empty{}, nil
}
