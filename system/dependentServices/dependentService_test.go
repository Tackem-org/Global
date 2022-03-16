package dependentServices_test

import (
	"fmt"
	"testing"

	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/stretchr/testify/assert"
)

func TestDependentServiceAddress(t *testing.T) {
	r := &dependentServices.DependentService{
		URL:  "",
		Port: 0,
	}
	assert.Equal(t, r.URL, r.Address())
	rs := &dependentServices.DependentService{
		URL:  "127.0.0.1",
		Port: 50000,
	}
	assert.Equal(t, fmt.Sprintf("%s:%d", rs.URL, rs.Port), rs.Address())
}
