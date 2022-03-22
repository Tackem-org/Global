package sysErrors_test

import (
	"testing"

	"github.com/Tackem-org/Global/sysErrors"
	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	assert.Error(t, sysErrors.MasterDownError)
	assert.Error(t, sysErrors.ServiceDownError)
	assert.Error(t, sysErrors.ServiceInactiveError)
	assert.Error(t, sysErrors.ConfigTypeError)
	assert.Error(t, sysErrors.ConfigValueError)
}
