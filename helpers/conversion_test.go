package helpers_test

import (
	"testing"

	"github.com/Tackem-org/Global/helpers"
	"github.com/stretchr/testify/assert"
)

func TestIntSliceToStringSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected []string
	}{
		{
			input:    []int{},
			expected: []string{},
		},
		{
			input:    []int{1},
			expected: []string{"1"},
		},
		{
			input:    []int{1, 2},
			expected: []string{"1", "2"},
		},
	}

	for _, test := range tests {
		data := helpers.IntSliceToStringSlice(test.input)
		assert.Equal(t, test.expected, data)
	}

}

func TestStringSliceToIntSlice(t *testing.T) {
	tests := []struct {
		input    []string
		expected []int
	}{
		{
			input:    []string{},
			expected: []int{},
		},
		{
			input:    []string{"1"},
			expected: []int{1},
		},
		{
			input:    []string{"1", "2"},
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		data := helpers.StringSliceToIntSlice(test.input)
		assert.Equal(t, test.expected, data)
	}

}

func TestInterfaceSliceToStringSlice(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected []string
	}{
		{
			input:    []interface{}{},
			expected: []string{},
		},
		{
			input:    []interface{}{1},
			expected: []string{"1"},
		},
		{
			input:    []interface{}{1, 2},
			expected: []string{"1", "2"},
		},
	}

	for _, test := range tests {
		data := helpers.InterfaceSliceToStringSlice(test.input)
		assert.Equal(t, test.expected, data)
	}

}

func TestInterfaceSliceToIntSlice(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected []int
		pass     bool
	}{
		{
			input:    []interface{}{},
			expected: []int{},
			pass:     true,
		},
		{
			input:    []interface{}{"1"},
			expected: []int{1},
			pass:     true,
		},
		{
			input:    []interface{}{"1", "2"},
			expected: []int{1, 2},
			pass:     true,
		},
		{
			input: []interface{}{"string"},
			pass:  false,
		},
		{
			input: []interface{}{"2.2"},
			pass:  false,
		},
	}

	for _, test := range tests {
		data, err := helpers.InterfaceSliceToIntSlice(test.input)
		if test.pass {
			assert.Equal(t, test.expected, data)
		} else {
			assert.Nil(t, data)
			assert.NotNil(t, err)
		}
	}

}

func TestStringToIntSlice(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
		pass     bool
	}{
		{
			input:    "1",
			expected: []int{1},
			pass:     true,
		},
		{
			input:    "1,2",
			expected: []int{1, 2},
			pass:     true,
		},
		{
			input: "",
			pass:  false,
		},
		{
			input: "string",
			pass:  false,
		},
		{
			input: "2.2",
			pass:  false,
		},
		{
			input: "string,1",
			pass:  false,
		},
	}

	for _, test := range tests {
		data, err := helpers.StringToIntSlice(test.input)
		if test.pass {
			assert.Equal(t, test.expected, data)
		} else {
			assert.Nil(t, data)
			assert.NotNil(t, err)
		}
	}

}

func TestStringToStringSlice(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
		pass     bool
	}{
		{
			input:    "1",
			expected: []string{"1"},
			pass:     true,
		},
		{
			input:    "1,2",
			expected: []string{"1", "2"},
			pass:     true,
		},
		{
			input:    "test",
			expected: []string{"test"},
			pass:     true,
		},
		{
			input:    "test,another",
			expected: []string{"test", "another"},
			pass:     true,
		},
		{
			input:    "",
			expected: []string{""},
		},
		{
			input:    "string",
			expected: []string{"string"},
		},
		{
			input:    "2.2",
			expected: []string{"2.2"},
		},
		{
			input:    "string,1",
			expected: []string{"string", "1"},
		},
	}

	for _, test := range tests {
		data := helpers.StringToStringSlice(test.input)
		assert.Equal(t, test.expected, data)
	}

}

//TODO WORKING DOWN FROM HERE AND TEST ALL THEN LOCKER THEN WORK DOWN
// will need to expand tests once I have ints converted in helper
func TestStringToStringMap(t *testing.T) {

	tests := []struct {
		input    string
		expected map[string]interface{}
		pass     bool
	}{
		{
			input:    `{"Test": 1.1, "Test2": "test"}`,
			expected: map[string]interface{}{"Test": 1.1, "Test2": "test"},
			pass:     true,
		},
		{
			input: `"Test": 1.1, "Test2": "test"}`,
			pass:  false,
		},
		{
			input: `[{"Test": 1.1, "Test2": "test"},{"Test": 2.1, "Test2": "test2"}]`,
			pass:  false,
		},
		{
			input: `[1.1, "test", 2.1, "test2"]`,
			pass:  false,
		},
	}

	for _, test := range tests {
		data, err := helpers.StringToStringMap(test.input)
		if test.pass {
			assert.Nil(t, err)
			assert.Equal(t, test.expected, data)
		} else {
			assert.Empty(t, data)
			assert.NotNil(t, err)
		}
	}

}

// func TestStringToDuration(t *testing.T) {
// 	data, err := helpers.StringToDuration("")

// }

// func TestDurationToString(t *testing.T) {
// 	data := helpers.DurationToString(time.Duration(100))

// }

// func TestMapStringInterfaceToMapStringString(t *testing.T) {
// 	data := helpers.MapStringInterfaceToMapStringString(map[string]interface{}{})

// }
