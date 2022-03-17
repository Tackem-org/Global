package helpers_test

import (
	"net"
	"testing"

	"github.com/Tackem-org/Global/helpers"
	"github.com/stretchr/testify/assert"
)

func MockLookupIP(host string) ([]net.IP, error) {
	l := net.ParseIP("127.0.0.1")
	b := net.ParseIP("")
	if host == "localhost" {
		return []net.IP{
			l,
		}, nil
	} else if host == "unknown" {
		return []net.IP{
			b,
		}, nil
	}
	return []net.IP{
		l,
	}, nil
}
func TestGrabIPFromURL(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{in: "localhost", expected: "127.0.0.1"},
		{in: "127.0.0.1", expected: "127.0.0.1"},
		{in: "172.0.0.1", expected: "172.0.0.1"},
		{in: "192.168.0.1", expected: "192.168.0.1"},
		{in: "unknown", expected: "<nil>"},
	}
	helpers.LookUPIP = MockLookupIP
	for _, test := range tests {
		assert.Equal(t, test.expected, helpers.GrabIPFromURL(test.in), test.in)
	}
}
