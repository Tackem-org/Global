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
		{
			part:  "{{test4}}{{second}}",
			pass:  false,
			parts: 0,
		},
		{
			part:  "?{{test5}}",
			pass:  false,
			parts: 0,
		},
		{
			part:  "{{test6}}?",
			pass:  false,
			parts: 0,
		},
		{
			part:  "{{test7:test7:test7}}",
			pass:  false,
			parts: 0,
		},
		{
			part:  "{{test8}}",
			pass:  false,
			parts: 0,
		},
		{
			part:  "{{test9:test9}}",
			pass:  false,
			parts: 0,
		},
		{
			part:  "{{number:test10-}}",
			pass:  false,
			parts: 0,
		},
		{
			part:  "{{string:test10-}}",
			pass:  false,
			parts: 0,
		},
		{
			part:  "{{string:test11}}",
			pass:  true,
			parts: 2,
		},
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
	tests := []struct {
		path string
		pass bool
	}{
		{
			path: "",
			pass: true,
		},
		{
			path: "/",
			pass: true,
		},
		{
			path: "/{{number:test}}",
			pass: true,
		},
		{
			path: "/{{test}}",
			pass: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.pass, helpers.CheckPath(test.path), test.path)
	}
}
