package helpers_test

import (
	"testing"

	"github.com/Tackem-org/Global/helpers"
	"github.com/stretchr/testify/assert"
)

func TestRandStr(t *testing.T) {
	for length := 1; length <= 16; length++ {
		for i := 0; i < 100; i++ {
			data := helpers.RandStr(length)
			assert.NotEmpty(t, data)
			assert.Equal(t, length, len(data))
		}
	}
}
