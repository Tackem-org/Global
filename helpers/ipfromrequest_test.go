package helpers_test

import (
	"net/http"
	"testing"

	"github.com/Tackem-org/Global/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGetIP(t *testing.T) {
	httpRequest1, _ := http.NewRequest("GET", "", nil)
	httpRequest1.Header.Set("X-FORWARDED-FOR", "127.0.0.1:50000")
	assert.Equal(t, "127.0.0.1", helpers.GetIP(httpRequest1))

	httpRequest2, _ := http.NewRequest("GET", "", nil)
	httpRequest2.RemoteAddr = "127.0.0.1:50000"
	assert.Equal(t, "127.0.0.1", helpers.GetIP(httpRequest2))
}
