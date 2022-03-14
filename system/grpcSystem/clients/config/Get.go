package config

import (
	"github.com/Tackem-org/Global/system/grpcSystem/connections"
	pb "github.com/Tackem-org/Proto/pb/config"

	"google.golang.org/grpc"
)

func (cc *ConfigClient) Get(request *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
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
