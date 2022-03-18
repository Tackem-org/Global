package regClient_test

import (
	"testing"

	"github.com/Tackem-org/Global/system/grpcSystem/servers/regClient"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
	pb "github.com/Tackem-org/Proto/pb/regclient"
	"github.com/stretchr/testify/assert"
)

func TestRequiredDown(t *testing.T) {
	s := regClient.RegClientServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.RequiredDown(ctx1, &pb.RequiredDownRequest{})
	assert.NotNil(t, r1)
	assert.Nil(t, err1)
	assert.False(t, r1.Success)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}
	r := &requiredServices.RequiredService{
		BaseID: "Testd1",
	}
	requiredServices.Add(r)

	ctx2 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r2, err2 := s.RequiredDown(ctx2, &pb.RequiredDownRequest{BaseId: r.BaseID})
	assert.NotNil(t, r2)
	assert.Nil(t, err2)
	assert.True(t, r2.Success)

	ctx3 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r3, err3 := s.RequiredDown(ctx3, &pb.RequiredDownRequest{BaseId: r.BaseID})
	assert.NotNil(t, r3)
	assert.Nil(t, err3)
	assert.False(t, r3.Success)
	assert.Equal(t, "Required Service Not Found", r3.ErrorMessage)

	requiredServices.Remove(r.BaseID)
}
