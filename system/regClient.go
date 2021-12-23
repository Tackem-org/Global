package system

import (
	"context"
	"sync"

	"github.com/Tackem-org/Global/logging"
	pb "github.com/Tackem-org/Proto/pb/regclient"
	"google.golang.org/grpc/metadata"
)

type RegClientServer struct {
	mu sync.RWMutex
	pb.UnimplementedRegClientServer
}

func NewRegClientServer() *RegClientServer {
	return &RegClientServer{}
}

func (r *RegClientServer) MasterGoingDown(ctx context.Context, in *pb.MasterGoingDownRequest) (*pb.MasterGoingDownResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	ok, err := r.checkKey(ctx)
	if ok {
		//TODO NEED TO DEAL WITH REASON HERE
		logging.Info("Master Is Down")
		MUp.Down()
	}
	return &pb.MasterGoingDownResponse{
		Success:      ok,
		ErrorMessage: err,
	}, nil

}

func (r *RegClientServer) MasterBackUp(ctx context.Context, in *pb.MasterBackUpRequest) (*pb.MasterBackUpResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	ok, err := r.checkKey(ctx)
	if ok {
		logging.Info("Master Is Back Up")
		MUp.Up()
	}
	return &pb.MasterBackUpResponse{
		Success:      ok,
		ErrorMessage: err,
	}, nil

}

func (r *RegClientServer) checkKey(ctx context.Context) (bool, string) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, "error retrieving header"
	}
	baseIDvalues := md.Get("baseID")
	if len(baseIDvalues) != 1 {
		return false, "baseID not found in header"
	}
	baseID := baseIDvalues[0]
	if baseID == "" {
		return false, "base id is blank"
	}
	keyvalues := md.Get("key")
	if len(keyvalues) != 1 {
		return false, "key not found in header"
	}
	key := keyvalues[0]
	if key == "" {
		return false, "key is blank"
	}

	if regData.GetBaseID() == baseID && regData.GetKey() == key {
		return true, ""
	}
	return false, "values not correct"
}
