package remoteWeb

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	pb "github.com/Tackem-org/Global/pb/remoteweb"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
)

func (r *RemoteWebServer) WebSocket(ctx context.Context, in *pb.WebSocketRequest) (*pb.WebSocketResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.WebSocketResponse{StatusCode: http.StatusInternalServerError, HideErrorFromUser: true, ErrorMessage: err}, nil
	}
	d, _ := helpers.StringToStringMap([]byte(in.DataJson))

	webSocketRequest := structs.SocketRequest{
		Command: in.Command,
		User:    structs.GetUserData(in.GetUser()),
		Data:    d,
	}

	command := strings.Split(in.Command, ".")

	var response *structs.SocketReturn
	var err error
	if command[0] == "panel" {
		Panel := setupData.Data.GetPanel(command[1])
		if Panel == nil {
			return &pb.WebSocketResponse{
				StatusCode:   http.StatusNotFound,
				ErrorMessage: "web socket panel not found",
			}, nil
		}
		response, err = Panel.SocketCall(&webSocketRequest)
	} else {
		socketItem := setupData.Data.GetSocket(in.Command)
		if socketItem == nil {
			return &pb.WebSocketResponse{
				StatusCode:   http.StatusNotFound,
				ErrorMessage: "web socket not found",
			}, nil
		}
		response, err = socketItem.Call(&webSocketRequest)
	}

	if err != nil {
		logging.Error("[GRPC Remote Web Socket Request] %s:%s", in.Command, err.Error())
		return &pb.WebSocketResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "error with the system",
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
