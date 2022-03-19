package remoteWeb_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/remoteWeb"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
	"github.com/stretchr/testify/assert"
)

func TestWebSocket(t *testing.T) {
	s := remoteWeb.RemoteWebServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.WebSocket(ctx1, &pb.WebSocketRequest{})
	assert.Nil(t, err1)
	assert.Equal(t, uint32(http.StatusInternalServerError), r1.StatusCode)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}

	setupData.Data = &setupData.SetupData{
		ServiceType: "service",
		StaticFS:    &MockEmbed{},
		Sockets: []*setupData.SocketItem{
			{
				Command:           "test",
				Permission:        "",
				AdminOnly:         false,
				RequiredVariables: []string{},
				Call: func(in *structs.SocketRequest) (*structs.SocketReturn, error) {
					return &structs.SocketReturn{
						StatusCode:   http.StatusOK,
						ErrorMessage: "",
						TellAll:      false,
						Data:         map[string]interface{}{},
					}, nil
				},
			},
			{
				Command:           "fail",
				Permission:        "",
				AdminOnly:         false,
				RequiredVariables: []string{},
				Call: func(in *structs.SocketRequest) (*structs.SocketReturn, error) {
					return nil, errors.New("FAIL")
				},
			},
		},
	}

	ctx2 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r2, err2 := s.WebSocket(ctx2, &pb.WebSocketRequest{Command: "NewTest", User: &pb.UserData{}})
	assert.Nil(t, err2)
	assert.Equal(t, uint32(http.StatusNotFound), r2.StatusCode)

	ctx3 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r3, err3 := s.WebSocket(ctx3, &pb.WebSocketRequest{Command: "fail", User: &pb.UserData{}})
	assert.Nil(t, err3)
	assert.Equal(t, uint32(http.StatusInternalServerError), r3.StatusCode)

	ctxPass := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	rPass, errPass := s.WebSocket(ctxPass, &pb.WebSocketRequest{Command: "test", User: &pb.UserData{}})
	assert.Nil(t, errPass)
	assert.Equal(t, uint32(http.StatusOK), rPass.StatusCode)
}
