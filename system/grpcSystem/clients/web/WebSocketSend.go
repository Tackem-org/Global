package web

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"
	pb "github.com/Tackem-org/Proto/pb/web"
	"google.golang.org/grpc"
)

func WebSocketSend(request *pb.SendWebSocketRequest) bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.client.web.WebSocketSend")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] request=%+v", request)
	conn, err := connections.Master()
	if err != nil {
		logging.Error("[Web Socket Send] Cannot Connect to Master: %s", err.Error())
		return false
	}
	defer conn.Close()
	client := pb.NewWebClient(conn)
	header, ctx, cancel := connections.MasterHeader()
	defer cancel()
	if _, err := client.SendWebSocket(ctx, request, grpc.Header(&header)); err != nil {
		logging.Error("[Web Socket Send] Error with the Server: %s", err.Error())
		return false
	}
	return true
}
