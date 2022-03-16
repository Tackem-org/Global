package web_test

import (
	"context"
	"errors"
	"log"
	"net"
	"testing"

	"github.com/Tackem-org/Global/system/grpcSystem/clients/web"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"

	pb "github.com/Tackem-org/Proto/pb/web"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type MockWebServer struct {
	pb.UnimplementedWebServer
}

func (c *MockWebServer) SendTask(ctx context.Context, in *pb.TaskMessage) (*pb.SendTaskResponse, error) {
	if in.Task == "FAIL" {
		return nil, errors.New("FAIL")
	}
	return &pb.SendTaskResponse{}, nil
}

func (c *MockWebServer) RemoveTask(ctx context.Context, in *pb.RemoveTaskRequest) (*pb.RemoveTaskResponse, error) {
	if in.Task == "FAIL" {
		return nil, errors.New("FAIL")
	}
	return &pb.RemoveTaskResponse{}, nil
}

func (c *MockWebServer) SendWebSocket(ctx context.Context, in *pb.SendWebSocketRequest) (*pb.SendWebSocketResponse, error) {
	if in.Command == "FAIL" {
		return nil, errors.New("FAIL")
	}
	return &pb.SendWebSocketResponse{}, nil
}

func startGRPCServer() (*grpc.Server, *bufconn.Listener) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)
	srv := grpc.NewServer()

	pb.RegisterWebServer(srv, &MockWebServer{})

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
func TestWebServerAddTask(t *testing.T) {
	wc := web.WebClient{}
	assert.False(t, wc.AddTask(&pb.TaskMessage{}))

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	assert.True(t, wc.AddTask(&pb.TaskMessage{}))
	assert.False(t, wc.AddTask(&pb.TaskMessage{Task: "FAIL"}))
}

func TestWebServerRemoveTask(t *testing.T) {
	wc := web.WebClient{}
	assert.False(t, wc.RemoveTask(&pb.RemoveTaskRequest{}))

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	assert.True(t, wc.RemoveTask(&pb.RemoveTaskRequest{}))
	assert.False(t, wc.RemoveTask(&pb.RemoveTaskRequest{Task: "FAIL"}))
}

func TestWebServerWebSocketSend(t *testing.T) {
	wc := web.WebClient{}
	assert.False(t, wc.WebSocketSend(&pb.SendWebSocketRequest{}))

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	assert.True(t, wc.WebSocketSend(&pb.SendWebSocketRequest{}))
	assert.False(t, wc.WebSocketSend(&pb.SendWebSocketRequest{Command: "FAIL"}))
}
