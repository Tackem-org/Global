package web_test

import (
	"testing"

	"github.com/Tackem-org/Global/system/grpcSystem/clients/web"
	pb "github.com/Tackem-org/Proto/pb/web"
	"github.com/stretchr/testify/assert"
)

type MockWebClient struct{}

func (wc *MockWebClient) AddTask(request *pb.TaskMessage) bool {
	return true
}

func (wc *MockWebClient) RemoveTask(request *pb.RemoveTaskRequest) bool {
	return true
}

func (mwc *MockWebClient) WebSocketSend(request *pb.SendWebSocketRequest) bool {
	return true
}

func TestAddTask(t *testing.T) {
	assert.False(t, web.AddTask(&pb.TaskMessage{}))
	web.I = &MockWebClient{}
	assert.True(t, web.AddTask(&pb.TaskMessage{}))
	web.I = nil
}

func TestRemoveTask(t *testing.T) {
	assert.False(t, web.RemoveTask(&pb.RemoveTaskRequest{}))
	web.I = &MockWebClient{}
	assert.True(t, web.RemoveTask(&pb.RemoveTaskRequest{}))
	web.I = nil
}

func TestWebSocketSend(t *testing.T) {
	assert.False(t, web.WebSocketSend(&pb.SendWebSocketRequest{}))
	web.I = &MockWebClient{}
	assert.True(t, web.WebSocketSend(&pb.SendWebSocketRequest{}))
	web.I = nil
}
