package web

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"
	"github.com/Tackem-org/Global/system/grpcSystem/headers"
	pb "github.com/Tackem-org/Proto/pb/web"
	"google.golang.org/grpc"
)

func RemoveTask(request *pb.RemoveTaskRequest) bool {
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
