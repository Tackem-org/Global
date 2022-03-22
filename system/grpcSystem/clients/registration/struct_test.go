package registration_test

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/Tackem-org/Global/pb/registration"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/registration"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type MockRegistrationServer struct {
	pb.UnimplementedRegistrationServer
}

func (rc *MockRegistrationServer) Activate(ctx context.Context, in *pb.ActivateRequest) (*pb.ActivateResponse, error) {
	return &pb.ActivateResponse{}, nil
}

func (rc *MockRegistrationServer) Deactivate(ctx context.Context, in *pb.DeactivateRequest) (*pb.DeactivateResponse, error) {
	return &pb.DeactivateResponse{}, nil
}

func (rc *MockRegistrationServer) Deregister(ctx context.Context, in *pb.DeregisterRequest) (*pb.DeregisterResponse, error) {
	return &pb.DeregisterResponse{}, nil
}

func (rc *MockRegistrationServer) Disconnect(ctx context.Context, in *pb.DisconnectRequest) (*pb.DisconnectResponse, error) {
	return &pb.DisconnectResponse{}, nil
}

func (rc *MockRegistrationServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{}, nil
}

func startGRPCServer() (*grpc.Server, *bufconn.Listener) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)
	srv := grpc.NewServer()

	pb.RegisterRegistrationServer(srv, &MockRegistrationServer{})

	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()

	getBufDialer := func(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
		return func(ctx context.Context, url string) (net.Conn, error) {
			return listener.Dial()
		}
	}

	connections.ExtraOptions = append(connections.ExtraOptions, grpc.WithContextDialer(getBufDialer(listener)))
	return srv, listener
}

func TestRegistrationActivate(t *testing.T) {
	rc := registration.RegistrationClient{}
	response1, err1 := rc.Activate(&pb.ActivateRequest{})
	assert.IsType(t, &pb.ActivateResponse{}, response1)
	assert.NotNil(t, err1)

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	response2, err2 := rc.Activate(&pb.ActivateRequest{})
	assert.IsType(t, &pb.ActivateResponse{}, response2)
	assert.Nil(t, err2)
}

func TestRegistrationDeactivate(t *testing.T) {
	rc := registration.RegistrationClient{}
	response1, err1 := rc.Deactivate(&pb.DeactivateRequest{})
	assert.IsType(t, &pb.DeactivateResponse{}, response1)
	assert.NotNil(t, err1)

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	response2, err2 := rc.Deactivate(&pb.DeactivateRequest{})
	assert.IsType(t, &pb.DeactivateResponse{}, response2)
	assert.Nil(t, err2)
}

func TestRegistrationDeregister(t *testing.T) {
	rc := registration.RegistrationClient{}
	response1, err1 := rc.Deregister(&pb.DeregisterRequest{})
	assert.IsType(t, &pb.DeregisterResponse{}, response1)
	assert.NotNil(t, err1)

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	response2, err2 := rc.Deregister(&pb.DeregisterRequest{})
	assert.IsType(t, &pb.DeregisterResponse{}, response2)
	assert.Nil(t, err2)
}

func TestRegistrationDisconnect(t *testing.T) {
	rc := registration.RegistrationClient{}
	response1, err1 := rc.Disconnect(&pb.DisconnectRequest{})
	assert.IsType(t, &pb.DisconnectResponse{}, response1)
	assert.NotNil(t, err1)

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	response2, err2 := rc.Disconnect(&pb.DisconnectRequest{})
	assert.IsType(t, &pb.DisconnectResponse{}, response2)
	assert.Nil(t, err2)
}

func TestRegistrationRegister(t *testing.T) {
	rc := registration.RegistrationClient{}
	response1, err1 := rc.Register(&pb.RegisterRequest{})
	assert.IsType(t, &pb.RegisterResponse{}, response1)
	assert.NotNil(t, err1)

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	response2, err2 := rc.Register(&pb.RegisterRequest{})
	assert.IsType(t, &pb.RegisterResponse{}, response2)
	assert.Nil(t, err2)
}
