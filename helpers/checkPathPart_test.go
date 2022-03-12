package helpers_test

import (
	"testing"

	"github.com/Tackem-org/Global/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCheckPathPart(t *testing.T) {
	tests := []struct {
		part  string
		pass  bool
		parts int
	}{
		{
			part:  "test1pass",
			pass:  true,
			parts: 0,
		},
		{
			part:  "{test2",
			pass:  false,
			parts: 0,
		},
		{
			part:  "{{test3}",
			pass:  false,
			parts: 0,
		},
	}

	for _, test := range tests {
		var parts int
		pass, b := helpers.CheckPathPart(test.part)
		assert.Equal(t, test.pass, pass)
		if b == nil {
			parts = 0
		} else {
			parts = len(b)
		}
		assert.Equal(t, test.parts, parts)
	}

}

// func TestCheckPath(t *testing.T) {
// }
