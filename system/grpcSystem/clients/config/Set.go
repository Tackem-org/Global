package config

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"
	pb "github.com/Tackem-org/Proto/pb/config"

	"google.golang.org/grpc"
)

func Set(request *pb.SetConfigRequest) (*pb.SetConfigResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.client.config.Set")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] request=%+v", request)
	conn, err := connections.Master()
	if err != nil {
		return &pb.SetConfigResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}
	defer conn.Close()
	client := pb.NewConfigClient(conn)
	header, ctx, cancel := connections.MasterHeader()
	defer cancel()
	return client.Set(ctx, request, grpc.Header(&header))
}
