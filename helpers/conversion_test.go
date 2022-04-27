package helpers_test

import (
	"testing"
	"time"

	"github.com/Tackem-org/Global/helpers"
	"github.com/stretchr/testify/assert"
)

func TestIntSliceToStringSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected []string
	}{
		{[]int{}, []string{}},
		{[]int{1}, []string{"1"}},
		{[]int{1, 2}, []string{"1", "2"}},
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
		{[]string{}, []int{}},
		{[]string{"1"}, []int{1}},
		{[]string{"1", "2"}, []int{1, 2}},
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
		{[]interface{}{}, []string{}},
		{[]interface{}{1}, []string{"1"}},
		{[]interface{}{1, 2}, []string{"1", "2"}},
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
		{[]interface{}{}, []int{}, true},
		{[]interface{}{"1"}, []int{1}, true},
		{[]interface{}{"1", "2"}, []int{1, 2}, true},
		{[]interface{}{int(1), int8(2), int16(3), int32(4), int64(5), uint(6), uint8(7), uint16(8), uint32(9), uint64(10), float32(11.1), float64(12.1), "13"}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, true},
		{input: []interface{}{"string"}, pass: false},
		{input: []interface{}{"2.2"}, pass: false},
		{input: []interface{}{false}, pass: false},
	}

	for i, test := range tests {
		data, err := helpers.InterfaceSliceToIntSlice(test.input)
		if test.pass {
			assert.Equal(t, test.expected, data)
		} else {
			assert.Nil(t, data)
			assert.NotNil(t, err, i)
		}
	}

}

func TestStringToIntSlice(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
		pass     bool
	}{
		{"1", []int{1}, true},
		{"1,2", []int{1, 2}, true},
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
	}{
		{"1", []string{"1"}},
		{"1,2", []string{"1", "2"}},
		{"test", []string{"test"}},
		{"test,another", []string{"test", "another"}},
		{"", []string{""}},
		{"string", []string{"string"}},
		{"2.2", []string{"2.2"}},
		{"string,1", []string{"string", "1"}},
	}

	for _, test := range tests {
		data := helpers.StringToStringSlice(test.input)
		assert.Equal(t, test.expected, data)
	}

}

func TestStringToStringMap(t *testing.T) {

	tests := []struct {
		input    string
		expected map[string]interface{}
		pass     bool
	}{
		{
			input:    `{"int":1, "float":1.1, "string": "test"}`,
			expected: map[string]interface{}{"int": 1, "float": 1.1, "string": "test"},
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
		{
			input:    `{"test1": {"test2": {"test3": "data"}}}`,
			expected: map[string]interface{}{"test1": map[string]interface{}{"test2": map[string]interface{}{"test3": "data"}}},
			pass:     true,
		},
	}

	for _, test := range tests {
		data, err := helpers.StringToStringMap([]byte(test.input))
		if test.pass {
			assert.Nil(t, err)
			assert.Equal(t, test.expected, data)
		} else {
			assert.Empty(t, data)
			assert.NotNil(t, err)
		}
	}

}

func TestStringToDuration(t *testing.T) {

	tests := []struct {
		input    string
		expected time.Duration
		pass     bool
	}{
		{"5.9", time.Duration(5), true},
		{"105", time.Duration(105), true},
		{"1ns", time.Duration(time.Nanosecond), true},
		{"1us", time.Duration(time.Microsecond), true},
		{"1ms", time.Duration(time.Millisecond), true},
		{"1s", time.Duration(time.Second), true},
		{"1m", time.Duration(time.Minute), true},
		{"1h", time.Duration(time.Hour), true},
		{"1d", time.Duration(24 * time.Hour), true},
		{"1w", time.Duration(7 * 24 * time.Hour), true},
		{
			input: "friday",
			pass:  false,
		},
	}

	for _, test := range tests {
		data, err := helpers.StringToDuration(test.input)
		if test.pass {
			assert.Nil(t, err)
			assert.Equal(t, test.expected, data)
		} else {
			assert.Empty(t, data)
			assert.NotNil(t, err)
		}
	}

}

func TestDurationToString(t *testing.T) {
	tests := []struct {
		input    time.Duration
		expected string
	}{
		{time.Duration(time.Nanosecond), "1ns"},
		{time.Duration(time.Microsecond), "1µs"},
		{time.Duration(time.Millisecond), "1ms"},
		{time.Duration(time.Second), "1s"},
		{time.Duration(time.Minute), "1m0s"},
		{time.Duration(time.Hour), "1h0m0s"},
		{time.Duration(24 * time.Hour), "1d0h0m0s"},
		{time.Duration(7 * 24 * time.Hour), "1w0d0h0m0s"},
	}
	for _, test := range tests {
		data := helpers.DurationToString(test.input)
		assert.Equal(t, test.expected, data)
	}

}

func TestMapStringInterfaceToMapStringString(t *testing.T) {
	tests := []struct {
		input    map[string]interface{}
		expected map[string]string
	}{
		{
			map[string]interface{}{}, map[string]string{},
		},
		{
			map[string]interface{}{"a": "text", "b": 1, "c": 2.2}, map[string]string{"a": "text", "b": "1", "c": "2.2"},
		},
	}

	for _, test := range tests {
		data := helpers.MapStringInterfaceToMapStringString(test.input)
		assert.Equal(t, test.expected, data)
	}

}
