package sysErrors_test

import (
	"testing"

	"github.com/Tackem-org/Global/sysErrors"
	"github.com/stretchr/testify/assert"
)

func TestMasterDownError(t *testing.T) {
	sd := sysErrors.MasterDownError{}
	assert.Error(t, &sd)
	assert.Equal(t, "master is down ", sd.Error())
}

func TestSystemDownError(t *testing.T) {
	sd := sysErrors.ServiceDownError{}
	assert.Error(t, &sd)
	assert.Equal(t, "service is down ", sd.Error())
}

func TestSystemInactiveError(t *testing.T) {
	si := sysErrors.ServiceInactiveError{}
	assert.Error(t, &si)
	assert.Equal(t, "service is inactive ", si.Error())
}
