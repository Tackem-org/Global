package web

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"
	pb "github.com/Tackem-org/Proto/pb/web"
	"google.golang.org/grpc"
)

func AddTask(request *pb.TaskMessage) bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.client.web.AddTask")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] request=%+v", request)
	conn, err := connections.Master()
	if err != nil {
		logging.Errorf("[Add Task] Cannot Connect to Master: %s", err.Error())
		return false
	}
	defer conn.Close()
	client := pb.NewWebClient(conn)
	header, ctx, cancel := connections.MasterHeader()
	defer cancel()
	if _, err := client.SendTask(ctx, request, grpc.Header(&header)); err != nil {
		logging.Errorf("[Add Task] Error with the Server: %s", err.Error())
		return false
	}
	return true
}
