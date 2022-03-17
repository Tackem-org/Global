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

func TestConnectionInfostruct(t *testing.T) {
	key := "Test"
	IP := "127.0.0.1"
	a := masterData.ConnectionInfostruct{Key: key, IP: IP}
	assert.True(t, a.CheckKey(key))
	assert.False(t, a.CheckKey("AAA"))
	assert.True(t, a.CheckIP(IP))
	assert.False(t, a.CheckIP("192.168.0.1"))
}
