package remoteWeb_test

import (
	"testing"

	pb "github.com/Tackem-org/Global/pb/remoteweb"
	pbw "github.com/Tackem-org/Global/pb/web"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/remoteWeb"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
	"github.com/stretchr/testify/assert"
)

var SendNotifications bool

func TestNotifications(t *testing.T) {
	s := remoteWeb.RemoteWebServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.Notifications(ctx1, &pb.NotificationsRequest{})
	assert.Nil(t, err1)
	assert.False(t, r1.Success)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}

	setupData.Data = &setupData.SetupData{
		ServiceType:         "service",
		NotificationGrabber: func() []*pbw.NotificationMessage { return []*pbw.NotificationMessage{} },
	}

	ctxPass := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	rPass, errPass := s.Notifications(ctxPass, &pb.NotificationsRequest{})
	assert.Nil(t, errPass)
	assert.True(t, rPass.Success)
}
