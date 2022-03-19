package remoteWeb_test

import (
	"testing"

	"github.com/Tackem-org/Global/system/grpcSystem/servers/remoteWeb"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
	"github.com/Tackem-org/Proto/pb/web"
	"github.com/stretchr/testify/assert"
)

var SendTasks bool

func TestTasks(t *testing.T) {
	s := remoteWeb.RemoteWebServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.Tasks(ctx1, &pb.TasksRequest{})
	assert.Nil(t, err1)
	assert.False(t, r1.Success)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}

	setupData.Data = &setupData.SetupData{
		ServiceType: "service",
		TaskGrabber: func() []*web.TaskMessage { return []*web.TaskMessage{} },
	}

	ctxPass := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	rPass, errPass := s.Tasks(ctxPass, &pb.TasksRequest{})
	assert.Nil(t, errPass)
	assert.True(t, rPass.Success)
}
