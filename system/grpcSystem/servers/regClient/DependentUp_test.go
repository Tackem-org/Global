package regClient_test

import (
	"testing"

	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/regClient"
	"github.com/Tackem-org/Global/system/masterData"
	pb "github.com/Tackem-org/Proto/pb/regclient"
	"github.com/stretchr/testify/assert"
)

func TestDependentUp(t *testing.T) {
	s := regClient.RegClientServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.DependentUp(ctx1, &pb.DependentUpRequest{})
	assert.NotNil(t, r1)
	assert.Nil(t, err1)
	assert.False(t, r1.Success)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}
	d := &dependentServices.DependentService{
		BaseID: "Testd1",
	}
	d.UP.Down()
	dependentServices.Add(d)

	ctx2 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r2, err2 := s.DependentUp(ctx2, &pb.DependentUpRequest{BaseId: d.BaseID})
	assert.NotNil(t, r2)
	assert.Nil(t, err2)
	assert.True(t, r2.Success)

	ctx3 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r3, err3 := s.DependentUp(ctx3, &pb.DependentUpRequest{BaseId: d.BaseID})
	assert.NotNil(t, r3)
	assert.Nil(t, err3)
	assert.False(t, r3.Success)

	dependentServices.Remove(d.BaseID)
}
