package web_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/web"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"

	pb "github.com/Tackem-org/Global/pb/web"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type MockLogging struct{}

func (l *MockLogging) Setup(logFile string, verbose bool)                          {}
func (l *MockLogging) Shutdown()                                                   {}
func (l *MockLogging) CustomLogger(prefix string) *log.Logger                      { return log.New(nil, prefix+": ", 0) }
func (l *MockLogging) Custom(prefix string, message string, values ...interface{}) {}
func (l *MockLogging) Info(message string, values ...interface{})                  {}
func (l *MockLogging) Warning(message string, values ...interface{})               {}
func (l *MockLogging) Error(message string, values ...interface{})                 {}
func (l *MockLogging) Todo(message string, values ...interface{})                  {}
func (l *MockLogging) Fatal(message string, values ...interface{}) error {
	return fmt.Errorf(message, values...)
}

type MockWebServer struct {
	pb.UnimplementedWebServer
}

func (c *MockWebServer) SendTask(ctx context.Context, in *pb.TaskMessage) (*pb.SendTaskResponse, error) {
	if in.Task == "FAIL" {
		return nil, errors.New("FAIL")
	}
	return &pb.SendTaskResponse{}, nil
}

func (c *MockWebServer) SendNotification(ctx context.Context, in *pb.NotificationMessage) (*pb.SendNotificationResponse, error) {
	if in.Notification == "FAIL" {
		return nil, errors.New("FAIL")
	}
	return &pb.SendNotificationResponse{}, nil
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
	logging.I = &MockLogging{}
	assert.False(t, web.AddTask(&pb.TaskMessage{}))

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	assert.True(t, web.AddTask(&pb.TaskMessage{}))
	assert.False(t, web.AddTask(&pb.TaskMessage{Task: "FAIL"}))
}

func TestWebServerAddNotification(t *testing.T) {
	logging.I = &MockLogging{}
	assert.False(t, web.AddNotification(&pb.NotificationMessage{}))

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	assert.True(t, web.AddNotification(&pb.NotificationMessage{}))
	assert.False(t, web.AddNotification(&pb.NotificationMessage{Notification: "FAIL"}))
}

func TestWebServerRemoveTask(t *testing.T) {
	logging.I = &MockLogging{}
	assert.False(t, web.RemoveTask(&pb.RemoveTaskRequest{}))

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	assert.True(t, web.RemoveTask(&pb.RemoveTaskRequest{}))
	assert.False(t, web.RemoveTask(&pb.RemoveTaskRequest{Task: "FAIL"}))
}

func TestWebServerWebSocketSend(t *testing.T) {
	logging.I = &MockLogging{}
	assert.False(t, web.WebSocketSend(&pb.SendWebSocketRequest{}))

	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer srv.Stop()

	assert.True(t, web.WebSocketSend(&pb.SendWebSocketRequest{}))
	assert.False(t, web.WebSocketSend(&pb.SendWebSocketRequest{Command: "FAIL"}))
}
