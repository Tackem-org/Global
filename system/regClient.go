package system

import (
	"context"
	"sync"

	"github.com/Tackem-org/Global/channels"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	pb "github.com/Tackem-org/Proto/pb/regclient"
	"google.golang.org/grpc/metadata"
)

type RegClientServer struct {
	mu sync.Mutex
	pb.UnimplementedRegClientServer
}

func (r *RegClientServer) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS|debug.GRPCSERVER, "CALLED:[system.(r *RegClientServer) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error)] {in=%v}", in)
	r.mu.Lock()
	defer r.mu.Unlock()
	logging.Info("Health Check OK")
	return &pb.HealthCheckResponse{
		Healthy: healthcheckHealthy,
		Issues:  healthcheckIssues,
	}, nil
}

func (r *RegClientServer) AddDependant(ctx context.Context, in *pb.AddDependantRequest) (*pb.AddDependantResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS|debug.GRPCSERVER, "CALLED:[system.(r *RegClientServer) AddDependant(ctx context.Context, in *pb.AddDependantRequest) (*pb.AddDependantResponse, error)] {in=%v}", in)
	for _, s := range requiredServices {
		if s.BaseID == in.BaseId {
			return &pb.AddDependantResponse{
				Success: true,
			}, nil
		}
	}
	dependentServices = append(dependentServices, DependentService{
		BaseID:    in.BaseId,
		Key:       in.Key,
		IPAddress: in.IpAddress,
		Port:      in.Port,
		SingleRun: in.SingleRun,
	})
	return &pb.AddDependantResponse{
		Success: true,
	}, nil
}

func (r *RegClientServer) RemoveDependant(ctx context.Context, in *pb.RemoveDependantRequest) (*pb.RemoveDependantResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS|debug.GRPCSERVER, "CALLED:[system.(r *RegClientServer) RemoveDependant(ctx context.Context, in *pb.RemoveDependantRequest) (*pb.RemoveDependantResponse, error)] {in=%v}", in)
	for index, s := range requiredServices {
		if s.BaseID == in.BaseId {
			dependentServices = append(dependentServices[:index], dependentServices[index+1:]...)
			return &pb.RemoveDependantResponse{
				Success: true,
			}, nil
		}
	}
	return &pb.RemoveDependantResponse{
		Success:      false,
		ErrorMessage: "Service Not Found",
	}, nil
}

func (r *RegClientServer) MasterGoingDown(ctx context.Context, in *pb.MasterGoingDownRequest) (*pb.MasterGoingDownResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS|debug.GRPCSERVER, "CALLED:[system.(r *RegClientServer) MasterGoingDown(ctx context.Context, in *pb.MasterGoingDownRequest) (*pb.MasterGoingDownResponse, error)] {in=%v}", in)
	r.mu.Lock()
	defer r.mu.Unlock()

	ok, err := r.checkKey(ctx)
	if ok {
		switch in.GetReason() {
		case pb.MasterGoingDownReason_FullShutdown:
			logging.Info("Master Is Down, Shutting down this service")
			channels.Root.Shutdown <- true

		case pb.MasterGoingDownReason_Shutdown:
			logging.Info("Master Is Down")
			MUp.Down()
		}
	}
	return &pb.MasterGoingDownResponse{
		Success:      ok,
		ErrorMessage: err,
	}, nil
}

func (r *RegClientServer) MasterBackUp(ctx context.Context, in *pb.MasterBackUpRequest) (*pb.MasterBackUpResponse, error) {
	logging.Debugf(debug.FUNCTIONCALLS|debug.GRPCSERVER, "CALLED:[system.(r *RegClientServer) MasterBackUp(ctx context.Context, in *pb.MasterBackUpRequest) (*pb.MasterBackUpResponse, error)] {in=%v}", in)
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
		Active:       Active,
	}, nil
}

func (r *RegClientServer) checkKey(ctx context.Context) (bool, string) {
	logging.Debug(debug.FUNCTIONCALLS|debug.GRPCSERVER, "CALLED:[system.(r *RegClientServer) checkKey(ctx context.Context) (bool, string)]")
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
