package regClient_test

import (
	"testing"

	"github.com/Tackem-org/Global/logging"
	pb "github.com/Tackem-org/Global/pb/regclient"
	"github.com/Tackem-org/Global/system/channels"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/regClient"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/stretchr/testify/assert"
)

func TestShutdownService(t *testing.T) {
	logging.I = &MockLogging{}
	s := regClient.RegClientServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.ShutdownService(ctx1, &pb.ShutdownServiceRequest{})
	assert.NotNil(t, r1)
	assert.Nil(t, err1)
	assert.False(t, r1.Success)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}

	channels.Setup()

	ctx2 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r2, err2 := s.ShutdownService(ctx2, &pb.ShutdownServiceRequest{})
	assert.NotNil(t, r2)
	assert.Nil(t, err2)
	assert.True(t, r2.Success)
	assert.True(t, <-channels.Root.Shutdown)

}
