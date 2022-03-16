package requiredServices_test

import (
	"fmt"
	"testing"

	"github.com/Tackem-org/Global/system/requiredServices"
	"github.com/stretchr/testify/assert"
)

func TestRequiredServiceAddress(t *testing.T) {
	r := &requiredServices.RequiredService{
		URL:  "",
		Port: 0,
	}
	assert.Equal(t, r.URL, r.Address())
	rs := &requiredServices.RequiredService{
		URL:  "127.0.0.1",
		Port: 50000,
	}
	assert.Equal(t, fmt.Sprintf("%s:%d", rs.URL, rs.Port), rs.Address())
}
