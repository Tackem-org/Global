package config

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"
	pb "github.com/Tackem-org/Proto/pb/config"

	"google.golang.org/grpc"
)

func Get(request *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.client.config.Get")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] request=%+v", request)
	conn, err := connections.Master()
	if err != nil {
		return &pb.GetConfigResponse{
			Key:   "",
			Value: "",
			Type:  pb.ValueType_failed,
		}, err
	}
	defer conn.Close()
	client := pb.NewConfigClient(conn)
	header, ctx, cancel := connections.MasterHeader()
	defer cancel()
	return client.Get(ctx, request, grpc.Header(&header))
}
