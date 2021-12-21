package regClientServer

import (
	"context"
	"sync"

	"github.com/Tackem-org/Global/system"
	pb "github.com/Tackem-org/Proto/pb/regclient"
)

type RegClientServer struct {
	mu sync.RWMutex
	pb.UnimplementedRegClientServer
}

func NewServer() *RegClientServer {
	return &RegClientServer{}
}

func (r *RegClientServer) MasterGoingDown(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	system.MasterUp = false
	return &pb.Empty{}, nil
}
