package debug_test

import (
	"testing"

	"github.com/Tackem-org/Global/logging/debug"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	tests := []struct {
		actual   debug.Mask
		set      debug.Mask
		expected debug.Mask
	}{
		{
			actual:   debug.FUNCTIONCALLS,
			set:      debug.FUNCTIONARGS,
			expected: debug.FUNCTIONCALLS | debug.FUNCTIONARGS,
		},
	}

	for i, test := range tests {
		test.actual.Set(test.set)
		assert.Equal(t, test.expected, test.actual, i)
	}
}

func TestClear(t *testing.T) {
	tests := []struct {
		actual   debug.Mask
		clear    debug.Mask
		expected debug.Mask
	}{
		{
			actual:   debug.FUNCTIONCALLS | debug.FUNCTIONARGS,
			clear:    debug.FUNCTIONARGS,
			expected: debug.FUNCTIONCALLS,
		},
		{
			actual:   debug.NONE,
			clear:    debug.FUNCTIONARGS,
			expected: debug.NONE,
		},
	}

	for i, test := range tests {
		test.actual.Clear(test.clear)
		assert.Equal(t, test.expected, test.actual, i)
	}
}

func TestToggle(t *testing.T) {
	tests := []struct {
		actual   debug.Mask
		toggle   debug.Mask
		expected debug.Mask
	}{
		{
			actual:   debug.FUNCTIONCALLS | debug.FUNCTIONARGS,
			toggle:   debug.FUNCTIONARGS,
			expected: debug.FUNCTIONCALLS,
		},
		{
			actual:   debug.NONE,
			toggle:   debug.FUNCTIONARGS,
			expected: debug.FUNCTIONARGS,
		},
	}

	for i, test := range tests {
		test.actual.Toggle(test.toggle)
		assert.Equal(t, test.expected, test.actual, i)
	}
}

func TestHas(t *testing.T) {
	tests := []struct {
		actual   debug.Mask
		has      debug.Mask
		expected bool
	}{
		{
			actual:   debug.FUNCTIONCALLS | debug.FUNCTIONARGS,
			has:      debug.FUNCTIONARGS,
			expected: true,
		},
		{
			actual:   debug.FUNCTIONCALLS | debug.FUNCTIONARGS,
			has:      debug.FUNCTIONCALLS | debug.FUNCTIONARGS,
			expected: true,
		},
		{
			actual:   debug.FUNCTIONARGS,
			has:      debug.FUNCTIONCALLS | debug.FUNCTIONARGS,
			expected: false,
		},
		{
			actual:   debug.NONE,
			has:      debug.FUNCTIONARGS,
			expected: false,
		},
	}

	for i, test := range tests {
		assert.Equal(t, test.expected, test.actual.Has(test.has), i)
	}
}

func TestHasAny(t *testing.T) {
	tests := []struct {
		actual   debug.Mask
		hasAny   debug.Mask
		expected bool
	}{
		{
			actual:   debug.FUNCTIONCALLS | debug.FUNCTIONARGS,
			hasAny:   debug.FUNCTIONARGS,
			expected: true,
		},
		{
			actual:   debug.FUNCTIONCALLS | debug.FUNCTIONARGS,
			hasAny:   debug.FUNCTIONCALLS | debug.FUNCTIONARGS,
			expected: true,
		},
		{
			actual:   debug.FUNCTIONARGS,
			hasAny:   debug.FUNCTIONCALLS | debug.FUNCTIONARGS,
			expected: true,
		},
		{
			actual:   debug.NONE,
			hasAny:   debug.FUNCTIONARGS,
			expected: false,
		},
	}

	for i, test := range tests {
		assert.Equal(t, test.expected, test.actual.HasAny(test.hasAny), i)
	}
}
