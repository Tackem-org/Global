package headers_test

import (
	"testing"

	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/grpcSystem/headers"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
	"github.com/Tackem-org/Global/system/setupData"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestHeader(t *testing.T) {
	baseID := "TestID"
	key := "TestKey"
	header, _, _ := headers.Header(map[string]string{"baseID": baseID, "key": key})
	assert.IsType(t, metadata.MD{}, header)
	assert.Equal(t, baseID, header.Get("baseID")[0])
	assert.Equal(t, key, header.Get("key")[0])
}

func TestMasterHeader(t *testing.T) {
	header, _, _ := headers.MasterHeader()
	assert.IsType(t, metadata.MD{}, header)
	assert.Equal(t, setupData.BaseID, header.Get("baseID")[0])
	assert.Equal(t, setupData.Key, header.Get("key")[0])
}

func TestRegistrationHeader(t *testing.T) {
	header, _, _ := headers.RegistrationHeader()
	assert.IsType(t, metadata.MD{}, header)
	assert.Equal(t, masterData.Info.RegistrationKey, header.Get("registrationkey")[0])
}

func TestRequiredServiceHeader(t *testing.T) {
	r := &requiredServices.RequiredService{Key: "TestRequired"}
	header, _, _ := headers.RequiredServiceHeader(r)
	assert.IsType(t, metadata.MD{}, header)
	assert.Equal(t, setupData.BaseID, header.Get("baseID")[0])
	assert.Equal(t, r.Key, header.Get("key")[0])
}

func TestDependentServiceHeader(t *testing.T) {
	d := &dependentServices.DependentService{Key: "TestDependent"}
	header, _, _ := headers.DependentServiceHeader(d)
	assert.IsType(t, metadata.MD{}, header)
	assert.Equal(t, setupData.BaseID, header.Get("baseID")[0])
	assert.Equal(t, d.Key, header.Get("key")[0])
}
