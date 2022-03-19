package remoteWeb_test

import (
	"net/http"
	"testing"

	"github.com/Tackem-org/Global/system/grpcSystem/servers/remoteWeb"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	s := remoteWeb.RemoteWebServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.File(ctx1, &pb.FileRequest{})
	assert.Nil(t, err1)
	assert.Equal(t, uint32(http.StatusInternalServerError), r1.StatusCode)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}

	setupData.Data = &setupData.SetupData{
		ServiceType: "service",
		StaticFS:    &MockEmbed{Files: map[string]*MockFile{"testfile": {Data: "test"}}},
	}

	ctx2 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r2, err2 := s.File(ctx2, &pb.FileRequest{Path: "fail/"})
	assert.Nil(t, err2)
	assert.Equal(t, uint32(http.StatusForbidden), r2.StatusCode)

	ctx3 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r3, err3 := s.File(ctx3, &pb.FileRequest{Path: "fail"})
	assert.Nil(t, err3)
	assert.Equal(t, uint32(http.StatusNotFound), r3.StatusCode)

	ctx4 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r4, err4 := s.File(ctx4, &pb.FileRequest{Path: "ISE"})
	assert.Nil(t, err4)
	assert.Equal(t, uint32(http.StatusInternalServerError), r4.StatusCode)

	ctxPass := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	rPass, errPass := s.File(ctxPass, &pb.FileRequest{Path: "testfile"})
	assert.Nil(t, errPass)
	assert.Equal(t, uint32(http.StatusOK), rPass.StatusCode)
}
