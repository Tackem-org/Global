package registration_test

import (
	"testing"

	pb "github.com/Tackem-org/Global/pb/registration"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/registration"
	"github.com/stretchr/testify/assert"
)

type MockRegistrationClient struct{}

func (mrc *MockRegistrationClient) Activate(request *pb.ActivateRequest) (*pb.ActivateResponse, error) {
	return &pb.ActivateResponse{}, nil
}
func (mrc *MockRegistrationClient) Deactivate(request *pb.DeactivateRequest) (*pb.DeactivateResponse, error) {
	return &pb.DeactivateResponse{}, nil
}
func (mrc *MockRegistrationClient) Deregister(request *pb.DeregisterRequest) (*pb.DeregisterResponse, error) {
	return &pb.DeregisterResponse{}, nil
}
func (mrc *MockRegistrationClient) Disconnect(request *pb.DisconnectRequest) (*pb.DisconnectResponse, error) {
	return &pb.DisconnectResponse{}, nil
}
func (mrc *MockRegistrationClient) Register(request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{}, nil
}

func TestActivate(t *testing.T) {
	registration.I = &MockRegistrationClient{}
	scr, err2 := registration.Activate(&pb.ActivateRequest{})
	assert.IsType(t, &pb.ActivateResponse{}, scr)
	assert.Nil(t, err2)
	registration.I = nil
}

func TestDeactivate(t *testing.T) {
	registration.I = &MockRegistrationClient{}
	scr, err2 := registration.Deactivate(&pb.DeactivateRequest{})
	assert.IsType(t, &pb.DeactivateResponse{}, scr)
	assert.Nil(t, err2)
	registration.I = nil
}

func TestDeregister(t *testing.T) {
	registration.I = &MockRegistrationClient{}
	scr, err2 := registration.Deregister(&pb.DeregisterRequest{})
	assert.IsType(t, &pb.DeregisterResponse{}, scr)
	assert.Nil(t, err2)
	registration.I = nil
}

func TestDisconnect(t *testing.T) {
	registration.I = &MockRegistrationClient{}
	scr, err2 := registration.Disconnect(&pb.DisconnectRequest{})
	assert.IsType(t, &pb.DisconnectResponse{}, scr)
	assert.Nil(t, err2)
	registration.I = nil
}

func TestRegister(t *testing.T) {
	registration.I = &MockRegistrationClient{}
	scr, err2 := registration.Register(&pb.RegisterRequest{})
	assert.IsType(t, &pb.RegisterResponse{}, scr)
	assert.Nil(t, err2)
	registration.I = nil
}
