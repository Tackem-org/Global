package helpers_test

import (
	"testing"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/stretchr/testify/assert"
)

func TestCheckPathPart(t *testing.T) {
	logging.I = &MockLogging{}
	tests := []struct {
		part  string
		pass  bool
		parts int
	}{
		{"test1pass", true, 0},
		{"{test2", false, 0},
		{"{{test3}", false, 0},
		{"{{test4}}{{second}}", false, 0},
		{"?{{test5}}", false, 0},
		{"{{test6}}?", false, 0},
		{"{{test7:test7:test7}}", false, 0},
		{"{{test8}}", false, 0},
		{"{{test9:test9}}", false, 0},
		{"{{number:test10-}}", false, 0},
		{"{{string:test10-}}", false, 0},
		{"{{string:test11}}", true, 2},
	}

	for _, test := range tests {
		var parts int
		pass, b := helpers.CheckPathPart(test.part)
		assert.Equal(t, test.pass, pass, test.part)
		if b == nil {
			parts = 0
		} else {
			parts = len(b)
		}
		assert.Equal(t, test.parts, parts, test.part)
	}

}

func TestCheckPath(t *testing.T) {
	logging.I = &MockLogging{}
	tests := []struct {
		path string
		pass bool
	}{
		{"", true},
		{"/", true},
		{"/{{number:test}}", true},
		{"/{{test}}", false},
	}

	for _, test := range tests {
		assert.Equal(t, test.pass, helpers.CheckPath(test.path), test.path)
	}
}
