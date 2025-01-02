package remoteWeb_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Tackem-org/Global/logging"
	pb "github.com/Tackem-org/Global/pb/remoteweb"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/remoteWeb"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
	"github.com/stretchr/testify/assert"
)

func TestPanel(t *testing.T) {
	logging.I = &MockLogging{}
	s := remoteWeb.RemoteWebServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.Panel(ctx1, &pb.PanelRequest{})
	assert.Nil(t, err1)
	assert.Equal(t, uint32(http.StatusInternalServerError), r1.StatusCode)
	assert.True(t, r1.HideErrorFromUser)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}

	setupData.Data = &setupData.SetupData{
		ServiceType: "service",
		Panels: []*setupData.PanelItem{
			{
				Name:        "test",
				Label:       "Test",
				Description: "Test Panel",
				Layout: setupData.PanelLayout{
					HorizontalAlign:       setupData.HCenter,
					VerticalAlign:         setupData.VCenter,
					Width:                 2,
					Height:                1,
					ScrollWidth:           false,
					ScrollHeight:          false,
					TitleBar:              true,
					CloseButton:           true,
					CloseFromOutsideClick: true,
					BackgroundOpacity:     50,
				},
				AdminOnly:  false,
				Permission: "",
				HTMLCall: func(in *structs.PanelRequest) (*structs.PanelReturn, error) {
					return &structs.PanelReturn{
						StatusCode: 200,
						PanelHTML:  "ok",
					}, nil
				},
			},
		},
	}

	ctx2 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r2, err2 := s.Panel(ctx2, &pb.PanelRequest{User: &pb.UserData{}})
	assert.Nil(t, err2)
	assert.Equal(t, uint32(http.StatusNotFound), r2.StatusCode)
	ctx3 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r3, err3 := s.Panel(ctx3, &pb.PanelRequest{Name: "test", User: &pb.UserData{}})

	assert.NotNil(t, r3)
	assert.Nil(t, err3)
	assert.Equal(t, http.StatusOK, int(r3.StatusCode))
}

func TestMakePanelResponse(t *testing.T) {
	logging.I = &MockLogging{}
	response1 := remoteWeb.MakePanelResponse(&pb.PanelRequest{}, &structs.PanelReturn{}, errors.New("TEST"))
	assert.Equal(t, http.StatusInternalServerError, int(response1.StatusCode))

	response2 := remoteWeb.MakePanelResponse(&pb.PanelRequest{}, &structs.PanelReturn{StatusCode: http.StatusOK}, nil)
	assert.Equal(t, http.StatusOK, int(response2.StatusCode))

}
