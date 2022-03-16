package masterData_test

import (
	"fmt"
	"testing"

	"github.com/Tackem-org/Global/system/masterData"
	"github.com/stretchr/testify/assert"
)

func TestInfostruct(t *testing.T) {
	r := &masterData.Infostruct{
		URL:  "",
		Port: 0,
	}
	assert.Equal(t, r.URL, r.Address())
	rs := &masterData.Infostruct{
		URL:  "127.0.0.1",
		Port: 50000,
	}
	assert.Equal(t, fmt.Sprintf("%s:%d", rs.URL, rs.Port), rs.Address())
}
