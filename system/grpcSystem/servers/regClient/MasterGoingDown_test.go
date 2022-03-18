package regClient_test

import (
	"testing"

	"github.com/Tackem-org/Global/channels"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/regClient"
	"github.com/Tackem-org/Global/system/masterData"
	pb "github.com/Tackem-org/Proto/pb/regclient"
	"github.com/stretchr/testify/assert"
)

func TestMasterGoingDown(t *testing.T) {
	s := regClient.RegClientServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.MasterGoingDown(ctx1, &pb.MasterGoingDownRequest{})
	assert.NotNil(t, r1)
	assert.Nil(t, err1)
	assert.False(t, r1.Success)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}

	channels.Setup()

	ctx2 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r2, err2 := s.MasterGoingDown(ctx2, &pb.MasterGoingDownRequest{Reason: pb.MasterGoingDownReason_FullShutdown})
	assert.NotNil(t, r2)
	assert.Nil(t, err2)
	assert.True(t, r2.Success)
	<-channels.Root.Shutdown

	ctx3 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r3, err3 := s.MasterGoingDown(ctx3, &pb.MasterGoingDownRequest{Reason: pb.MasterGoingDownReason_Shutdown})
	assert.NotNil(t, r3)
	assert.Nil(t, err3)
	assert.True(t, r3.Success)

	ctx4 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r4, err4 := s.MasterGoingDown(ctx4, &pb.MasterGoingDownRequest{Reason: pb.MasterGoingDownReason_Shutdown})
	assert.NotNil(t, r4)
	assert.Nil(t, err4)
	assert.False(t, r4.Success)

	ctx5 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r5, err5 := s.MasterGoingDown(ctx5, &pb.MasterGoingDownRequest{Reason: -1})
	assert.NotNil(t, r5)
	assert.Nil(t, err5)
	assert.False(t, r5.Success)

}
