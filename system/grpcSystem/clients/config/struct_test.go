package config_test

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/Tackem-org/Global/system/grpcSystem/clients/config"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/Tackem-org/Global/pb/config"
)

type MockConfigServer struct {
	pb.UnimplementedConfigServer
}

func (m *MockConfigServer) Set(ctx context.Context, in *pb.SetConfigRequest) (*pb.SetConfigResponse, error) {
	return &pb.SetConfigResponse{Success: true}, nil
}

func (c *MockConfigServer) Get(ctx context.Context, in *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	return &pb.GetConfigResponse{}, nil
}
func startGRPCServer() (*grpc.Server, *bufconn.Listener) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)
	srv := grpc.NewServer()

	pb.RegisterConfigServer(srv, &MockConfigServer{})

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

func TestConfigServerSet(t *testing.T) {

	cc := config.ConfigClient{}
	cr, err1 := cc.Set(&pb.SetConfigRequest{})
	assert.IsType(t, &pb.SetConfigResponse{}, cr)
	assert.False(t, cr.Success)
	assert.NotEmpty(t, cr.ErrorMessage)
	assert.NotNil(t, err1)

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	scr, err2 := cc.Set(&pb.SetConfigRequest{})
	assert.IsType(t, &pb.SetConfigResponse{}, scr)
	assert.True(t, scr.Success)
	assert.Empty(t, scr.ErrorMessage)
	assert.Nil(t, err2)
}

func TestConfigServerGet(t *testing.T) {

	cc := config.ConfigClient{}
	cr, err := cc.Get(&pb.GetConfigRequest{})
	assert.IsType(t, &pb.GetConfigResponse{}, cr)
	assert.NotNil(t, err)

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	scr, err2 := cc.Get(&pb.GetConfigRequest{})
	assert.IsType(t, &pb.GetConfigResponse{}, scr)
	assert.Nil(t, err2)
}
