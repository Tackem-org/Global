package config_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/Tackem-org/Global/config"
	"github.com/Tackem-org/Global/sysErrors"
	configClient "github.com/Tackem-org/Global/system/grpcSystem/clients/config"
	pb "github.com/Tackem-org/Proto/pb/config"
	"github.com/stretchr/testify/assert"
)

func TestGetBool(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test":  {Value: "true", Type: pb.ValueType_Bool},
		"bad":   {Value: "[][]", Type: pb.ValueType_Bool},
		"wrong": {Value: "1", Type: pb.ValueType_Int},
	}}

	value, err := config.GetBool("fail")
	assert.False(t, value)
	assert.Error(t, err)

	value, berr := config.GetBool("bad")
	assert.False(t, value)
	assert.Error(t, berr)
	assert.Error(t, berr, &sysErrors.ConfigValueError{})

	value, verr := config.GetBool("wrong")
	assert.False(t, value)
	assert.Error(t, verr)
	assert.Error(t, verr, &sysErrors.ConfigTypeError{})

	value, nilerr := config.GetBool("test")
	assert.True(t, value)
	assert.Nil(t, nilerr)
}

func TestGetFloat64(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test":  {Value: "1.1", Type: pb.ValueType_Float64},
		"bad":   {Value: "[][]", Type: pb.ValueType_Float64},
		"wrong": {Value: "test", Type: pb.ValueType_String},
	}}

	value, err := config.GetFloat64("fail")
	assert.Equal(t, 0.0, value)
	assert.Error(t, err)

	value, berr := config.GetFloat64("bad")
	assert.Equal(t, 0.0, value)
	assert.Error(t, berr)
	assert.Error(t, berr, &sysErrors.ConfigValueError{})

	value, verr := config.GetFloat64("wrong")
	assert.Equal(t, 0.0, value)
	assert.Error(t, verr)
	assert.Error(t, verr, &sysErrors.ConfigTypeError{})

	value, nilerr := config.GetFloat64("test")
	assert.Equal(t, 1.1, value)
	assert.Nil(t, nilerr)
}

func TestGetInt(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test":  {Value: "1", Type: pb.ValueType_Int},
		"bad":   {Value: "[][]", Type: pb.ValueType_Int},
		"wrong": {Value: "test", Type: pb.ValueType_String},
	}}

	value, err := config.GetInt("fail")
	assert.Equal(t, 0, value)
	assert.Error(t, err)

	value, berr := config.GetInt("bad")
	assert.Equal(t, 0, value)
	assert.Error(t, berr)
	assert.Error(t, berr, &sysErrors.ConfigValueError{})

	value, verr := config.GetInt("wrong")
	assert.Equal(t, 0, value)
	assert.Error(t, verr)
	assert.Error(t, verr, &sysErrors.ConfigTypeError{})

	value, nilerr := config.GetInt("test")
	assert.Equal(t, 1, value)
	assert.Nil(t, nilerr)
}

func TestGetInt32(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test":  {Value: "1", Type: pb.ValueType_Int32},
		"bad":   {Value: "[][]", Type: pb.ValueType_Int32},
		"wrong": {Value: "test", Type: pb.ValueType_String},
	}}

	value, err := config.GetInt32("fail")
	assert.Equal(t, int32(0), value)
	assert.Error(t, err)

	value, berr := config.GetInt32("bad")
	assert.Equal(t, int32(0), value)
	assert.Error(t, berr)
	assert.Error(t, berr, &sysErrors.ConfigValueError{})

	value, verr := config.GetInt32("wrong")
	assert.Equal(t, int32(0), value)
	assert.Error(t, verr)
	assert.Error(t, verr, &sysErrors.ConfigTypeError{})

	value, nilerr := config.GetInt32("test")
	assert.Equal(t, int32(1), value)
	assert.Nil(t, nilerr)
}

func TestGetInt64(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test":  {Value: "1", Type: pb.ValueType_Int64},
		"bad":   {Value: "[][]", Type: pb.ValueType_Int64},
		"wrong": {Value: "test", Type: pb.ValueType_String},
	}}

	value, err := config.GetInt64("fail")
	assert.Equal(t, int64(0), value)
	assert.Error(t, err)

	value, berr := config.GetInt64("bad")
	assert.Equal(t, int64(0), value)
	assert.Error(t, berr)
	assert.Error(t, berr, &sysErrors.ConfigValueError{})

	value, verr := config.GetInt64("wrong")
	assert.Equal(t, int64(0), value)
	assert.Error(t, verr)
	assert.Error(t, verr, &sysErrors.ConfigTypeError{})

	value, nilerr := config.GetInt64("test")
	assert.Equal(t, int64(1), value)
	assert.Nil(t, nilerr)
}

func TestGetUint(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test":  {Value: "1", Type: pb.ValueType_Uint},
		"bad":   {Value: "[][]", Type: pb.ValueType_Uint},
		"wrong": {Value: "test", Type: pb.ValueType_String},
	}}

	value, err := config.GetUint("fail")
	assert.Equal(t, uint(0), value)
	assert.Error(t, err)

	value, berr := config.GetUint("bad")
	assert.Equal(t, uint(0), value)
	assert.Error(t, berr)
	assert.Error(t, berr, &sysErrors.ConfigValueError{})

	value, verr := config.GetUint("wrong")
	assert.Equal(t, uint(0), value)
	assert.Error(t, verr)
	assert.Error(t, verr, &sysErrors.ConfigTypeError{})

	value, nilerr := config.GetUint("test")
	assert.Equal(t, uint(1), value)
	assert.Nil(t, nilerr)
}

func TestGetUint32(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test":  {Value: "1", Type: pb.ValueType_Uint32},
		"bad":   {Value: "[][]", Type: pb.ValueType_Uint32},
		"wrong": {Value: "test", Type: pb.ValueType_String},
	}}

	value, err := config.GetUint32("fail")
	assert.Equal(t, uint32(0), value)
	assert.Error(t, err)

	value, berr := config.GetUint32("bad")
	assert.Equal(t, uint32(0), value)
	assert.Error(t, berr)
	assert.Error(t, berr, &sysErrors.ConfigValueError{})

	value, verr := config.GetUint32("wrong")
	assert.Equal(t, uint32(0), value)
	assert.Error(t, verr)
	assert.Error(t, verr, &sysErrors.ConfigTypeError{})

	value, nilerr := config.GetUint32("test")
	assert.Equal(t, uint32(1), value)
	assert.Nil(t, nilerr)
}

func TestGetUint64(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test":  {Value: "1", Type: pb.ValueType_Uint64},
		"bad":   {Value: "[][]", Type: pb.ValueType_Uint64},
		"wrong": {Value: "test", Type: pb.ValueType_String},
	}}

	value, err := config.GetUint64("fail")
	assert.Equal(t, uint64(0), value)
	assert.Error(t, err)

	value, berr := config.GetUint64("bad")
	assert.Equal(t, uint64(0), value)
	assert.Error(t, berr)
	assert.Error(t, berr, &sysErrors.ConfigValueError{})

	value, verr := config.GetUint64("wrong")
	assert.Equal(t, uint64(0), value)
	assert.Error(t, verr)
	assert.Error(t, verr, &sysErrors.ConfigTypeError{})

	value, nilerr := config.GetUint64("test")
	assert.Equal(t, uint64(1), value)
	assert.Nil(t, nilerr)
}

func TestGetIntSlice(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test":  {Value: "1+2+3+4", Type: pb.ValueType_IntSlice},
		"bad":   {Value: "[][]", Type: pb.ValueType_IntSlice},
		"wrong": {Value: "test", Type: pb.ValueType_String},
	}}

	value, err := config.GetIntSlice("fail")
	assert.Equal(t, []int{}, value)
	assert.Error(t, err)

	value, berr := config.GetIntSlice("bad")
	assert.Equal(t, []int{}, value)
	assert.Error(t, berr)
	assert.Error(t, berr, &sysErrors.ConfigValueError{})

	value, verr := config.GetIntSlice("wrong")
	assert.Equal(t, []int{}, value)
	assert.Error(t, verr)
	assert.Error(t, verr, &sysErrors.ConfigTypeError{})

	value, nilerr := config.GetIntSlice("test")
	assert.Equal(t, []int{1, 2, 3, 4}, value)
	assert.Nil(t, nilerr)
}

func TestGetString(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test":  {Value: "test", Type: pb.ValueType_String},
		"wrong": {Value: "1", Type: pb.ValueType_Int},
	}}

	value, err := config.GetString("fail")
	assert.Equal(t, "", value)
	assert.Error(t, err)

	value, verr := config.GetString("wrong")
	assert.Equal(t, "", value)
	assert.Error(t, verr)
	assert.Error(t, verr, &sysErrors.ConfigTypeError{})

	value, nilerr := config.GetString("test")
	assert.Equal(t, "test", value)
	assert.Nil(t, nilerr)

}

func TestGetStringSlice(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test": {Value: "a,b,c,d", Type: pb.ValueType_StringSlice},

		"wrong": {Value: "1", Type: pb.ValueType_Int},
	}}

	value, err := config.GetStringSlice("fail")
	assert.Equal(t, []string{}, value)
	assert.Error(t, err)

	value, verr := config.GetStringSlice("wrong")
	assert.Equal(t, []string{}, value)
	assert.Error(t, verr)
	assert.Error(t, verr, &sysErrors.ConfigTypeError{})

	value, nilerr := config.GetStringSlice("test")
	assert.Equal(t, []string{"a", "b", "c", "d"}, value)
	assert.Nil(t, nilerr)
}

func TestGetTime(t *testing.T) {
	tn := time.Now().Round(time.Second)
	ts := strconv.FormatInt(tn.Unix(), 10)

	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test":  {Value: ts, Type: pb.ValueType_Time},
		"bad":   {Value: "[][]", Type: pb.ValueType_Time},
		"wrong": {Value: "test", Type: pb.ValueType_String},
	}}

	value, err := config.GetTime("fail")
	assert.Equal(t, time.Unix(0, 0), value)
	assert.Error(t, err)

	value, berr := config.GetTime("bad")
	assert.Equal(t, time.Unix(0, 0), value)
	assert.Error(t, berr)
	assert.Error(t, berr, &sysErrors.ConfigValueError{})

	value, verr := config.GetTime("wrong")
	assert.Equal(t, time.Unix(0, 0), value)
	assert.Error(t, verr)
	assert.Error(t, verr, &sysErrors.ConfigTypeError{})

	value, nilerr := config.GetTime("test")
	assert.Equal(t, tn, value)
	assert.Nil(t, nilerr)
}

func TestGetDuration(t *testing.T) {
	configClient.I = &MockConfig{Data: map[string]ConfigInfo{
		"test":  {Value: "1s", Type: pb.ValueType_Duration},
		"bad":   {Value: "[][]", Type: pb.ValueType_Duration},
		"wrong": {Value: "test", Type: pb.ValueType_String},
	}}

	value, err := config.GetDuration("fail")
	assert.Equal(t, time.Duration(0), value)
	assert.Error(t, err)

	value, berr := config.GetDuration("bad")
	assert.Equal(t, time.Duration(0), value)
	assert.Error(t, berr)
	assert.Error(t, berr, &sysErrors.ConfigValueError{})

	value, verr := config.GetDuration("wrong")
	assert.Equal(t, time.Duration(0), value)
	assert.Error(t, verr)
	assert.Error(t, verr, &sysErrors.ConfigTypeError{})

	value, nilerr := config.GetDuration("test")
	assert.Equal(t, time.Second, value)
	assert.Nil(t, nilerr)
}
