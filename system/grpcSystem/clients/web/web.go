package web

import (
	"github.com/Tackem-org/Global/logging"
	pb "github.com/Tackem-org/Global/pb/web"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"
	"github.com/Tackem-org/Global/system/grpcSystem/headers"
	"google.golang.org/grpc"
)

type WebClientInterface interface {
	AddTask(request *pb.TaskMessage) bool
	AddNotification(request *pb.NotificationMessage) bool
	RemoveTask(request *pb.RemoveTaskRequest) bool
	WebSocketSend(request *pb.SendWebSocketRequest) bool
}

var I WebClientInterface = &WebClient{}

func AddTask(request *pb.TaskMessage) bool {
	return I.AddTask(request)
}

func AddNotification(request *pb.NotificationMessage) bool {
	return I.AddNotification(request)
}

func RemoveTask(request *pb.RemoveTaskRequest) bool {
	return I.RemoveTask(request)
}

func WebSocketSend(request *pb.SendWebSocketRequest) bool {
	return I.WebSocketSend(request)
}

type WebClient struct{}

func (wc *WebClient) AddTask(request *pb.TaskMessage) bool {
	conn, err := connections.Master()
	if err != nil {
		logging.Error("[Add Task] Cannot Connect to Master: %s", err.Error())
		return false
	}
	defer conn.Close()
	client := pb.NewWebClient(conn)
	header, ctx, cancel := headers.MasterHeader()
	defer cancel()
	if _, err := client.SendTask(ctx, request, grpc.Header(&header)); err != nil {
		logging.Error("[Add Task] Error with the Server: %s", err.Error())
		return false
	}
	return true
}

func (wc *WebClient) AddNotification(request *pb.NotificationMessage) bool {
	conn, err := connections.Master()
	if err != nil {
		logging.Error("[Add Notification] Cannot Connect to Master: %s", err.Error())
		return false
	}
	defer conn.Close()
	client := pb.NewWebClient(conn)
	header, ctx, cancel := headers.MasterHeader()
	defer cancel()
	if _, err := client.SendNotification(ctx, request, grpc.Header(&header)); err != nil {
		logging.Error("[Add Notification] Error with the Server: %s", err.Error())
		return false
	}
	return true
}

func (wc *WebClient) RemoveTask(request *pb.RemoveTaskRequest) bool {
	conn, err := connections.Master()
	if err != nil {
		logging.Error("[Remove Task] Cannot Connect to Master: %s", err.Error())
		return false
	}
	defer conn.Close()
	client := pb.NewWebClient(conn)
	header, ctx, cancel := headers.MasterHeader()
	defer cancel()
	if _, err := client.RemoveTask(ctx, request, grpc.Header(&header)); err != nil {
		logging.Error("[Add Task] Error with the Server: %s", err.Error())
		return false
	}
	return true
}

func (wc *WebClient) WebSocketSend(request *pb.SendWebSocketRequest) bool {
	conn, err := connections.Master()
	if err != nil {
		logging.Error("[Web Socket Send] Cannot Connect to Master: %s", err.Error())
		return false
	}
	defer conn.Close()
	client := pb.NewWebClient(conn)
	header, ctx, cancel := headers.MasterHeader()
	defer cancel()
	if _, err := client.SendWebSocket(ctx, request, grpc.Header(&header)); err != nil {
		logging.Error("[Web Socket Send] Error with the Server: %s", err.Error())
		return false
	}
	return true
}
