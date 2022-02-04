package web

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Master/data"
	"github.com/Tackem-org/Master/web/authentication"

	pb "github.com/Tackem-org/Proto/pb/remoteweb"
	pbu "github.com/Tackem-org/Proto/pb/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/schollz/websocket"
	"github.com/schollz/websocket/wsjson"
)

func serviceWebsocketHandler(w http.ResponseWriter, r *http.Request, service *data.RegData) {
	logging.Debug(debug.FUNCTIONCALLS|debug.GPRCCLIENT, "CALLED:[web.serviceWebsocketHandler(w http.ResponseWriter, r *http.Request, service *registration.RegData)]")
	path := strings.Replace(r.URL.Path, Baseurl, "", 1)
	userID, _ := authentication.GetUserID(r)
	conn, err := service.GetGPRCConnection()
	if err != nil {
		errorPage(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	defer conn.Close()
	client := pb.NewRemoteWebClient(conn)

	ctxGPRC, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctxGPRC = metadata.NewOutgoingContext(ctxGPRC, metadata.MD{})

	request := &pb.ValidRequest{
		Path: path,
	}
	var response *pb.ValidResponse
	if strings.HasPrefix(path, "admin/") {
		response, _ = client.ValidWebSocket(ctxGPRC, request, grpc.Header(&metadata.MD{}))
	} else {
		response, _ = client.ValidAdminWebSocket(ctxGPRC, request, grpc.Header(&metadata.MD{}))
	}

	if !response.Found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		logging.Error(err.Error())
		return
	}

	ctxWS, cancel := context.WithTimeout(r.Context(), time.Hour*120000)
	defer cancel()

	for {
		var v map[string]interface{}
		err = wsjson.Read(ctxWS, c, &v)
		if err != nil {
			if websocket.CloseStatus(err) == websocket.StatusNoStatusRcvd {
				break
			}
			logging.Error(err.Error())
			c.Close(websocket.StatusInternalError, "internal error")
			break
		}

		command := v["command"].(string)
		if command == "close" || command == "quit" || command == "exit" {
			break
		}

		ctxGPRC, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		ctxGPRC = metadata.NewOutgoingContext(ctxGPRC, metadata.MD{})

		dataJson, _ := json.Marshal(v)
		request := &pb.WebSocketRequest{
			Path:     path,
			User:     getRemoteWebUserData(userID),
			DataJson: string(dataJson),
		}

		var response *pb.WebSocketResponse
		var errresp error
		if strings.HasPrefix(path, "admin/") {
			response, errresp = client.WebSocket(ctxGPRC, request, grpc.Header(&metadata.MD{}))
		} else {
			response, errresp = client.AdminWebSocket(ctxGPRC, request, grpc.Header(&metadata.MD{}))
		}
		if errresp != nil {
			errorPage(w, r, http.StatusInternalServerError, "Service Returned Bad Data")
			return
		}
		var w map[string]interface{}
		json.Unmarshal([]byte(response.DataJson), &w)
		wsjson.Write(ctxWS, c, map[string]interface{}{
			"status":        response.StatusCode,
			"error_message": response.ErrorMessage,
			"data":          w,
		})

	}

	if websocket.CloseStatus(err) == websocket.StatusGoingAway {
		err = nil
	}
	c.Close(websocket.StatusNormalClosure, "")

}

func getRemoteWebUserData(userID uint64) *pb.UserData {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.getUserData(userID) UserData]")
	errorData := &pb.UserData{Name: "ERROR", Initial: "", Icon: "", IsAdmin: false}
	userService := data.GetUserService()
	if userService == nil {
		return &pb.UserData{Name: "Service Down", Initial: "@", Icon: "", IsAdmin: true}
	}
	conn, err := userService.GetGPRCConnection()
	if err != nil {
		return errorData
	}
	defer conn.Close()

	client := pbu.NewUserClient(conn)

	ctx2, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx2 = metadata.NewOutgoingContext(ctx2, metadata.MD{})

	response, err := client.GetBaseData(ctx2, &pbu.GetBaseDataRequest{
		UserId: userID,
	}, grpc.Header(&metadata.MD{}))

	if err != nil || !response.Success {
		return errorData
	}
	return &pb.UserData{
		UserId:      response.UserId,
		Name:        response.Name,
		Initial:     response.Initial,
		Icon:        response.Icon,
		IsAdmin:     response.IsAdmin,
		Permissions: response.Permissions,
	}
}
