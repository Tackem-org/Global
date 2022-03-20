package sysErrors_test

import (
	"testing"

	"github.com/Tackem-org/Global/sysErrors"
	"github.com/stretchr/testify/assert"
)

func TestMasterDownError(t *testing.T) {
	e := sysErrors.MasterDownError{}
	assert.Error(t, &e)
	assert.Equal(t, "master is down ", e.Error())
}

func TestSystemDownError(t *testing.T) {
	e := sysErrors.ServiceDownError{}
	assert.Error(t, &e)
	assert.Equal(t, "service is down ", e.Error())
}

func TestSystemInactiveError(t *testing.T) {
	e := sysErrors.ServiceInactiveError{}
	assert.Error(t, &e)
	assert.Equal(t, "service is inactive ", e.Error())
}

func TestConfigTypeError(t *testing.T) {
	e := sysErrors.ConfigTypeError{}
	assert.Error(t, &e)
	assert.Equal(t, "config value type is wrong ", e.Error())
}

func TestConfigValueError(t *testing.T) {
	e := sysErrors.ConfigValueError{}
	assert.Error(t, &e)
	assert.Equal(t, "config value is bad ", e.Error())
}
