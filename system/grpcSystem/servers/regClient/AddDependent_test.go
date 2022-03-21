package regClient_test

import (
	"testing"

	pb "github.com/Tackem-org/Global/pb/regclient"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/regClient"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/stretchr/testify/assert"
)

func TestAddDependent(t *testing.T) {
	s := regClient.RegClientServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.AddDependent(ctx1, &pb.AddDependentRequest{})
	assert.NotNil(t, r1)
	assert.Nil(t, err1)
	assert.False(t, r1.Success)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}
	ctx2 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r2, err2 := s.AddDependent(ctx2, &pb.AddDependentRequest{})
	assert.NotNil(t, r2)
	assert.Nil(t, err2)
	assert.True(t, r2.Success)

	ctx3 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r3, err3 := s.AddDependent(ctx3, &pb.AddDependentRequest{})
	assert.NotNil(t, r3)
	assert.Nil(t, err3)
	assert.True(t, r3.Success)
}
