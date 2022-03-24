package web

import (
	pb "github.com/Tackem-org/Global/pb/web"
)

type WebClientInterface interface {
	AddTask(request *pb.TaskMessage) bool
	RemoveTask(request *pb.RemoveTaskRequest) bool
	WebSocketSend(request *pb.SendWebSocketRequest) bool
}

var I WebClientInterface = &WebClient{}

func AddTask(request *pb.TaskMessage) bool {
	return I.AddTask(request)
}

func RemoveTask(request *pb.RemoveTaskRequest) bool {
	return I.RemoveTask(request)
}

func WebSocketSend(request *pb.SendWebSocketRequest) bool {
	return I.WebSocketSend(request)
}
