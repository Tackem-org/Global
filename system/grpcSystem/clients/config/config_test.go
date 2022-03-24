package config_test

import (
	"testing"

	pb "github.com/Tackem-org/Global/pb/config"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/config"
	"github.com/stretchr/testify/assert"
)

type MockConfigClient struct{}

func (mcc *MockConfigClient) Set(request *pb.SetConfigRequest) (*pb.SetConfigResponse, error) {
	return &pb.SetConfigResponse{}, nil
}

func (mcc *MockConfigClient) Get(request *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	return &pb.GetConfigResponse{}, nil
}
func TestConfigGet(t *testing.T) {
	config.I = &MockConfigClient{}
	scr, err2 := config.Get(&pb.GetConfigRequest{})
	assert.IsType(t, &pb.GetConfigResponse{}, scr)
	assert.Nil(t, err2)
	config.I = nil
}

func TestConfigSet(t *testing.T) {
	config.I = &MockConfigClient{}
	scr, err2 := config.Set(&pb.SetConfigRequest{})
	assert.IsType(t, &pb.SetConfigResponse{}, scr)
	assert.Nil(t, err2)
	config.I = nil
}
