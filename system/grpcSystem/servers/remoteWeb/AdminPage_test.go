package remoteWeb_test

import (
	"net/http"
	"testing"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/remoteWeb"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
	"github.com/stretchr/testify/assert"
)

func TestAdminPage(t *testing.T) {
	s := remoteWeb.RemoteWebServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.AdminPage(ctx1, &pb.PageRequest{})
	assert.Nil(t, err1)
	assert.Equal(t, uint32(http.StatusInternalServerError), r1.StatusCode)
	assert.True(t, r1.HideErrorFromUser)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}

	setupData.Data = &setupData.SetupData{
		ServiceType: "service",
		AdminPaths: []*setupData.AdminPathItem{
			{
				Path: "/test",
				Call: func(in *structs.WebRequest) (*structs.WebReturn, error) {
					return &structs.WebReturn{
						StatusCode: 200,
						PageString: "ok",
					}, nil
				},
			},
		},
		StaticFS: &MockEmbed{},
	}

	ctx2 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r2, err2 := s.AdminPage(ctx2, &pb.PageRequest{User: &pb.UserData{IsAdmin: true}})
	assert.Nil(t, err2)
	assert.Equal(t, uint32(http.StatusNotFound), r2.StatusCode)
	ctx3 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r3, err3 := s.AdminPage(ctx3, &pb.PageRequest{BasePath: "/test", User: &pb.UserData{IsAdmin: true}})

	assert.NotNil(t, r3)
	assert.Nil(t, err3)
	assert.Equal(t, http.StatusOK, int(r3.StatusCode))
}
