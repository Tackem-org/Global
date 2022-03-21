package regClient_test

import (
	"testing"

	pb "github.com/Tackem-org/Global/pb/regclient"
	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/regClient"
	"github.com/Tackem-org/Global/system/healthCheck"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	s := regClient.RegClientServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.HealthCheck(ctx1, &pb.HealthCheckRequest{})
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
	assert.True(t, dependentServices.Add(d))

	ctx2 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r2, err2 := s.HealthCheck(ctx2, &pb.HealthCheckRequest{})
	assert.NotNil(t, r2)
	assert.Nil(t, err2)
	assert.True(t, r2.Success)
	assert.True(t, r2.Healthy)
	assert.Equal(t, 0, len(r2.Issues))

	healthCheck.SetIssue("Test Issue")
	ctx3 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r3, err3 := s.HealthCheck(ctx3, &pb.HealthCheckRequest{})
	assert.NotNil(t, r3)
	assert.Nil(t, err3)
	assert.True(t, r3.Success)
	assert.False(t, r3.Healthy)
	assert.Equal(t, 1, len(r3.Issues))
	dependentServices.Remove(d.BaseID)
}
