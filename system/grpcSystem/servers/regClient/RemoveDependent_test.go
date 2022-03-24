package regClient_test

import (
	"testing"

	"github.com/Tackem-org/Global/logging"
	pb "github.com/Tackem-org/Global/pb/regclient"
	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/regClient"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/stretchr/testify/assert"
)

func TestRemoveDependent(t *testing.T) {
	logging.I = &MockLogging{}
	s := regClient.RegClientServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.RemoveDependent(ctx1, &pb.RemoveDependentRequest{})
	assert.NotNil(t, r1)
	assert.Nil(t, err1)
	assert.False(t, r1.Success)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}
	d := &dependentServices.DependentService{
		BaseID: "Test1",
	}
	dependentServices.Add(d)
	ctx2 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r2, err2 := s.RemoveDependent(ctx2, &pb.RemoveDependentRequest{BaseId: d.BaseID})
	assert.NotNil(t, r2)
	assert.Nil(t, err2)
	assert.True(t, r2.Success)

	ctx3 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r3, err3 := s.RemoveDependent(ctx3, &pb.RemoveDependentRequest{BaseId: d.BaseID})
	assert.NotNil(t, r3)
	assert.Nil(t, err3)
	assert.False(t, r3.Success)
}
