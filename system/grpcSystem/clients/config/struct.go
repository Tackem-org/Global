package config

import (
	pb "github.com/Tackem-org/Global/pb/config"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"
	"github.com/Tackem-org/Global/system/grpcSystem/headers"

	"google.golang.org/grpc"
)

type ConfigClient struct{}

func (cc *ConfigClient) Set(request *pb.SetConfigRequest) (*pb.SetConfigResponse, error) {
	conn, err := connections.Master()
	if err != nil {
		return &pb.SetConfigResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}
	defer conn.Close()
	client := pb.NewConfigClient(conn)
	header, ctx, cancel := headers.MasterHeader()
	defer cancel()
	return client.Set(ctx, request, grpc.Header(&header))
}

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
	header, ctx, cancel := headers.MasterHeader()
	defer cancel()
	return client.Get(ctx, request, grpc.Header(&header))
}
