package dependentServices_test

import (
	"testing"

	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/stretchr/testify/assert"
)

func TestDependentServiceAddAndRemove(t *testing.T) {
	d := &dependentServices.DependentService{
		ServiceName: "Test1",
		ServiceType: "Test1",
		BaseID:      "1",
	}
	assert.True(t, dependentServices.Add(d), "Adding new service")
	assert.False(t, dependentServices.Add(d), "Adding the same service")
	assert.True(t, dependentServices.Remove(d.BaseID), "Removing the service")
	assert.False(t, dependentServices.Remove(d.BaseID), "Removing the service when already removed")
}

func TestDependentServiceUPAndDown(t *testing.T) {
	d := &dependentServices.DependentService{
		ServiceName: "Test2",
		ServiceType: "Test2",
		BaseID:      "2",
	}
	assert.True(t, dependentServices.Add(d), "Adding NEW Service")

	assert.False(t, dependentServices.Up(d.BaseID), "Trying to bring a service up that is already up")
	assert.True(t, dependentServices.Down(d.BaseID), "Taking down an up service")
	assert.False(t, dependentServices.Down(d.BaseID), "Trying to take down a service already down")
	assert.True(t, dependentServices.Up(d.BaseID), "Bring a Service Back Up")

	assert.True(t, dependentServices.Remove(d.BaseID), "Removing The Service")

	assert.False(t, dependentServices.Up(d.BaseID), "Trying to bring up a missing Service")
	assert.False(t, dependentServices.Down(d.BaseID), "Taking Down a missing Service")

}

func TestDependentServiceGetActive(t *testing.T) {
	d := &dependentServices.DependentService{
		ServiceName: "Test3",
		ServiceType: "Test3",
		BaseID:      "3",
	}

	count := len(dependentServices.GetActive())
	assert.True(t, dependentServices.Add(d), "Adding NEW Service")
	assert.Equal(t, count+1, len(dependentServices.GetActive()), "Check count of active")
	assert.True(t, dependentServices.Down(d.BaseID), "Taking down an up service")
	assert.Equal(t, count, len(dependentServices.GetActive()), "Check count of active")
	assert.True(t, dependentServices.Remove(d.BaseID), "Removing The Service")
}

func TestDependentServiceGetByBaseID(t *testing.T) {
	d := &dependentServices.DependentService{
		ServiceName: "Test4",
		ServiceType: "Test4",
		BaseID:      "4",
	}

	assert.Nil(t, dependentServices.GetByBaseID(d.BaseID))
	assert.True(t, dependentServices.Add(d), "Adding NEW Service")
	assert.NotNil(t, dependentServices.GetByBaseID(d.BaseID), "Check count of active")
	assert.True(t, dependentServices.Remove(d.BaseID), "Removing The Service")
	assert.Nil(t, dependentServices.GetByBaseID(d.BaseID), "Check count of active")
}
