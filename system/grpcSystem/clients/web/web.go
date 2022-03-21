package web

import (
	pb "github.com/Tackem-org/Global/pb/web"
)

type WebClientInterface interface {
	AddTask(request *pb.TaskMessage) bool
	RemoveTask(request *pb.RemoveTaskRequest) bool
	WebSocketSend(request *pb.SendWebSocketRequest) bool
}

var I WebClientInterface

func AddTask(request *pb.TaskMessage) bool {
	if I == nil {
		I = &WebClient{}
	}
	return I.AddTask(request)
}

func RemoveTask(request *pb.RemoveTaskRequest) bool {
	if I == nil {
		I = &WebClient{}
	}
	return I.RemoveTask(request)
}

func WebSocketSend(request *pb.SendWebSocketRequest) bool {
	if I == nil {
		I = &WebClient{}
	}
	return I.WebSocketSend(request)
}
