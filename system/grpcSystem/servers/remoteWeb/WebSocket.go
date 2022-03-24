package remoteWeb

import (
	"context"
	"encoding/json"
	"net/http"

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
	d, err := helpers.StringToStringMap(in.DataJson)

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
		logging.Error("[GRPC Remote Web Socket Request] %s:%s", in.Command, err.Error())
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
