package config_test

import (
	"errors"

	pb "github.com/Tackem-org/Global/pb/config"
)

type ConfigInfo struct {
	Value string
	Type  pb.ValueType
}
type MockConfig struct {
	Data map[string]ConfigInfo
}

func (mc *MockConfig) Get(request *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	if v, ok := mc.Data[request.Key]; ok {
		return &pb.GetConfigResponse{
			Key:   request.Key,
			Value: v.Value,
			Type:  v.Type,
		}, nil
	}
	return nil, errors.New("not found")
}

func (mc *MockConfig) Set(request *pb.SetConfigRequest) (*pb.SetConfigResponse, error) {
	if v, ok := mc.Data[request.Key]; ok {
		v.Value = request.Value
		mc.Data[request.Key] = v
		return &pb.SetConfigResponse{
			Success: true,
		}, nil
	}
	return &pb.SetConfigResponse{
		Success:      false,
		ErrorMessage: "Not Found",
	}, nil
}
