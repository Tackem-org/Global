package remoteWeb

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
)

func (r *RemoteWebServer) WebSocket(ctx context.Context, in *pb.WebSocketRequest) (*pb.WebSocketResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.servers.remoteWeb.RemoteWebServer{}.WebSocket")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] ctx in=%+v", in)
	var d map[string]interface{}
	json.Unmarshal([]byte(in.DataJson), &d)

	webSocketRequest := structs.SocketRequest{
		Command: in.Command,
		User:    structs.GetUserData(in.GetUser()),
		Data:    d,
	}

	socketItem := setupData.Data.GetSocket(in.Command)
	if socketItem == nil {
		return &pb.WebSocketResponse{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "Web Socket Not Found",
		}, nil
	}
	response, err := socketItem.Call(&webSocketRequest)

	if err != nil {
		logging.Errorf("[GRPC Remote Web Socket Request] %s:%s", in.Command, err.Error())
		return &pb.WebSocketResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "ERROR WITH THE SYSTEM",
		}, nil
	}

	returnJson, _ := json.Marshal(response.Data)
	return &pb.WebSocketResponse{
		StatusCode:   response.StatusCode,
		ErrorMessage: response.ErrorMessage,
		TellAll:      response.TellAll,
		DataJson:     string(returnJson),
	}, nil

}