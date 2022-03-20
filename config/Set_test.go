package config_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/Tackem-org/Global/config"
	configClient "github.com/Tackem-org/Global/system/grpcSystem/clients/config"
	pb "github.com/Tackem-org/Proto/pb/config"
	"github.com/stretchr/testify/assert"
)

func TestSetBool(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: "true", Type: pb.ValueType_Bool},
	}}

	pass, nilerr := config.SetBool("test", false)
	assert.True(t, pass)
	assert.Nil(t, nilerr)

	fail, err := config.SetBool("new", false)
	assert.False(t, fail)
	assert.NotNil(t, err)

}

func TestSetFloat64(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: "3.14", Type: pb.ValueType_Float64},
	}}

	pass, nilerr := config.SetFloat64("test", 3.14)
	assert.True(t, pass)
	assert.Nil(t, nilerr)

	fail, err := config.SetFloat64("new", 3.14)
	assert.False(t, fail)
	assert.NotNil(t, err)
}

func TestSetInt(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: "3", Type: pb.ValueType_Int},
	}}

	pass, nilerr := config.SetInt("test", 3)
	assert.True(t, pass)
	assert.Nil(t, nilerr)

	fail, err := config.SetInt("new", 3)
	assert.False(t, fail)
	assert.NotNil(t, err)
}

func TestSetInt32(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: "3", Type: pb.ValueType_Int32},
	}}

	pass, nilerr := config.SetInt32("test", 3)
	assert.True(t, pass)
	assert.Nil(t, nilerr)

	fail, err := config.SetInt32("new", 3)
	assert.False(t, fail)
	assert.NotNil(t, err)
}

func TestSetInt64(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: "3", Type: pb.ValueType_Int64},
	}}

	pass, nilerr := config.SetInt64("test", 3)
	assert.True(t, pass)
	assert.Nil(t, nilerr)

	fail, err := config.SetInt64("new", 3)
	assert.False(t, fail)
	assert.NotNil(t, err)
}

func TestSetUint(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: "3", Type: pb.ValueType_Uint},
	}}

	pass, nilerr := config.SetUint("test", 3)
	assert.True(t, pass)
	assert.Nil(t, nilerr)

	fail, err := config.SetUint("new", 3)
	assert.False(t, fail)
	assert.NotNil(t, err)
}

func TestSetUint32(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: "3", Type: pb.ValueType_Uint32},
	}}

	pass, nilerr := config.SetUint32("test", 3)
	assert.True(t, pass)
	assert.Nil(t, nilerr)

	fail, err := config.SetUint32("new", 3)
	assert.False(t, fail)
	assert.NotNil(t, err)
}

func TestSetUint64(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: "3", Type: pb.ValueType_Uint64},
	}}

	pass, nilerr := config.SetUint64("test", 3)
	assert.True(t, pass)
	assert.Nil(t, nilerr)

	fail, err := config.SetUint64("new", 3)
	assert.False(t, fail)
	assert.NotNil(t, err)
}

func TestSetIntSlice(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: "1+2+3+4", Type: pb.ValueType_IntSlice},
	}}

	pass, nilerr := config.SetIntSlice("test", []int{1, 2, 3, 4})
	assert.True(t, pass)
	assert.Nil(t, nilerr)

	fail, err := config.SetIntSlice("new", []int{1, 2, 3, 4})
	assert.False(t, fail)
	assert.NotNil(t, err)
}

func TestSetString(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: "test", Type: pb.ValueType_String},
	}}

	pass, nilerr := config.SetString("test", "Test")
	assert.True(t, pass)
	assert.Nil(t, nilerr)

	fail, err := config.SetString("new", "test")
	assert.False(t, fail)
	assert.NotNil(t, err)
}

func TestSetStringSlice(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: "a, b, c, d", Type: pb.ValueType_StringSlice},
	}}

	pass, nilerr := config.SetStringSlice("test", []string{"a", "b", "c", "d"})
	assert.True(t, pass)
	assert.Nil(t, nilerr)

	fail, err := config.SetStringSlice("new", []string{"a", "b", "c", "d"})
	assert.False(t, fail)
	assert.NotNil(t, err)
}

func TestSetTime(t *testing.T) {
	tn := time.Now()
	ts := strconv.FormatInt(tn.Unix(), 10)
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: ts, Type: pb.ValueType_Time},
	}}

	pass, nilerr := config.SetTime("test", tn)
	assert.True(t, pass)
	assert.Nil(t, nilerr)

	fail, err := config.SetTime("new", tn)
	assert.False(t, fail)
	assert.NotNil(t, err)
}

func TestSetDuration(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: "1s", Type: pb.ValueType_Duration},
	}}

	pass, nilerr := config.SetDuration("test", time.Second)
	assert.True(t, pass)
	assert.Nil(t, nilerr)

	fail, err := config.SetDuration("new", time.Second)
	assert.False(t, fail)
	assert.NotNil(t, err)
}
