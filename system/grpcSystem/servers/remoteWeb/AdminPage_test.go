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
	assert.NotNil(t, r1)
	assert.Nil(t, err1)
	assert.Equal(t, uint32(http.StatusInternalServerError), r1.StatusCode)
	assert.True(t, r1.HideErrorFromUser)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}

	//TODO NEED TO MAKE A TEMP ADMIN PATH TO RETURN DATA IN THIS AREA, THIS WILL BE SIMILAR FOR OTHERS IN THE PACKAGE
	setupData.Data = &setupData.SetupData{AdminPaths: []*setupData.AdminPathItem{{
		Path: "test",
		Call: func(in *structs.WebRequest) (*structs.WebReturn, error) {
			return &structs.WebReturn{
				StatusCode: 200,
				PageString: "ok",
			}, nil
		},
	}}}
	//TODO NEED TO MAKE A TEMP ADMIN PATH TO RETURN DATA IN THIS AREA, THIS WILL BE SIMILAR FOR OTHERS IN THE PACKAGE
	//TODO NEED TO MAKE A TEMP ADMIN PATH TO RETURN DATA IN THIS AREA, THIS WILL BE SIMILAR FOR OTHERS IN THE PACKAGE

	ctx2 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r2, err2 := s.AdminPage(ctx2, &pb.PageRequest{
		Path:     "",
		BasePath: "",
		User: &pb.UserData{
			UserId:      0,
			Name:        "",
			Icon:        "",
			IsAdmin:     false,
			Permissions: []string{},
		},
		Method:          "",
		QueryParamsJson: "",
		PostJson:        "",
		PathParamsJson:  "",
	})
	assert.NotNil(t, r2)
	assert.Nil(t, err2)
	assert.Equal(t, uint32(http.StatusNotFound), r2.StatusCode)

	ctx3 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r3, err3 := s.AdminPage(ctx3, &pb.PageRequest{
		Path:     "test",
		BasePath: "test",
		User: &pb.UserData{
			UserId:      0,
			Name:        "",
			Icon:        "",
			IsAdmin:     false,
			Permissions: []string{},
		},
		Method:          "",
		QueryParamsJson: "",
		PostJson:        "",
		PathParamsJson:  "",
	})
	assert.NotNil(t, r3)
	assert.Nil(t, err3)
	assert.Equal(t, uint32(http.StatusOK), r2.StatusCode)

}
