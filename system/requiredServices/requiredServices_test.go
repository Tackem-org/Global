package requiredServices_test

import (
	"testing"

	"github.com/Tackem-org/Global/system/requiredServices"
	"github.com/stretchr/testify/assert"
)

func TestRequiredServicesAddAndRemove(t *testing.T) {
	r := &requiredServices.RequiredService{
		ServiceName: "Test1",
		ServiceType: "Test1",
		BaseID:      "1",
	}
	assert.True(t, requiredServices.Add(r), "Adding new service")
	assert.False(t, requiredServices.Add(r), "Adding the same service")
	assert.True(t, requiredServices.Remove(r.BaseID), "Removing the service")
	assert.False(t, requiredServices.Remove(r.BaseID), "Removing the service when already removed")
}

func TestRequiredServicesUPAndDown(t *testing.T) {
	r := &requiredServices.RequiredService{
		ServiceName: "Test2",
		ServiceType: "Test2",
		BaseID:      "2",
	}
	assert.True(t, requiredServices.Add(r), "Adding NEW Service")

	assert.False(t, requiredServices.Up(r.BaseID), "Trying to bring a service up that is already up")
	assert.True(t, requiredServices.Down(r.BaseID), "Taking down an up service")
	assert.False(t, requiredServices.Down(r.BaseID), "Trying to take down a service already down")
	assert.True(t, requiredServices.Up(r.BaseID), "Bring a Service Back Up")

	assert.True(t, requiredServices.Remove(r.BaseID), "Removing The Service")

	assert.False(t, requiredServices.Up(r.BaseID), "Trying to bring up a missing Service")
	assert.False(t, requiredServices.Down(r.BaseID), "Taking Down a missing Service")

}

func TestRequiredServicesGet(t *testing.T) {
	r := &requiredServices.RequiredService{
		ServiceName: "Test3",
		ServiceType: "Test3",
		BaseID:      "3",
	}

	assert.Nil(t, requiredServices.Get(r.ServiceName, r.ServiceType), "Check count of active")
	assert.True(t, requiredServices.Add(r), "Adding NEW Service")
	assert.True(t, requiredServices.Down(r.BaseID), "Taking down an up service")
	assert.NotNil(t, requiredServices.Get(r.ServiceName, r.ServiceType), "Check count of active")
	assert.True(t, requiredServices.Remove(r.BaseID), "Removing The Service")
}

func TestRequiredServicesGetByBaseID(t *testing.T) {
	r := &requiredServices.RequiredService{
		ServiceName: "Test4",
		ServiceType: "Test4",
		BaseID:      "4",
	}

	assert.Nil(t, requiredServices.GetByBaseID(r.BaseID))
	assert.True(t, requiredServices.Add(r), "Adding NEW Service")
	assert.NotNil(t, requiredServices.GetByBaseID(r.BaseID), "Check count of active")
	assert.True(t, requiredServices.Remove(r.BaseID), "Removing The Service")
	assert.Nil(t, requiredServices.GetByBaseID(r.BaseID), "Check count of active")
}
