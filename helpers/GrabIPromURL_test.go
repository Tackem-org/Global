package helpers_test

import (
	"testing"

	"github.com/Tackem-org/Global/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGrabIPFromURL(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{in: "localhost", expected: "127.0.0.1"},
		{in: "127.0.0.1", expected: "127.0.0.1"},
		{in: "172.0.0.1", expected: "172.0.0.1"},
		{in: "192.168.0.1", expected: "192.168.0.1"},
		{in: "unknown", expected: ""},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, helpers.GrabIPFromURL(test.in), test.in)
	}
}
