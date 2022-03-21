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
	cr, err := config.Get(&pb.GetConfigRequest{})
	assert.IsType(t, &pb.GetConfigResponse{}, cr)
	assert.NotNil(t, err)

	config.I = &MockConfigClient{}
	scr, err2 := config.Get(&pb.GetConfigRequest{})
	assert.IsType(t, &pb.GetConfigResponse{}, scr)
	assert.Nil(t, err2)
	config.I = nil
}

func TestConfigSet(t *testing.T) {

	cr, err1 := config.Set(&pb.SetConfigRequest{})
	assert.IsType(t, &pb.SetConfigResponse{}, cr)
	assert.False(t, cr.Success)
	assert.NotEmpty(t, cr.ErrorMessage)
	assert.NotNil(t, err1)

	config.I = &MockConfigClient{}
	scr, err2 := config.Set(&pb.SetConfigRequest{})
	assert.IsType(t, &pb.SetConfigResponse{}, scr)
	assert.Nil(t, err2)
	config.I = nil
}
