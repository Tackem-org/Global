package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	configClient "github.com/Tackem-org/Global/system/grpcSystem/clients/config"
	pb "github.com/Tackem-org/Proto/pb/config"
)

func set(key string, value string) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.set")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%s", key, value)
	response, _ := configClient.Set(&pb.SetConfigRequest{Key: key, Value: value})
	return response.GetSuccess(), errors.New(response.GetErrorMessage())
}

func SetBool(key string, value bool) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetBool")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%t", key, value)
	return set(key, fmt.Sprintf("%t", value))
}

func SetFloat64(key string, value float64) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetFloat64")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%f", key, value)
	return set(key, fmt.Sprintf("%f", value))
}

func SetInt(key string, value int) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetInt")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%d", key, value)
	return set(key, fmt.Sprintf("%d", value))
}

func SetInt32(key string, value int32) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetInt32")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%d", key, value)
	return set(key, fmt.Sprintf("%d", value))
}

func SetInt64(key string, value int64) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetInt64")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%d", key, value)
	return set(key, fmt.Sprintf("%d", value))
}

func SetUint(key string, value uint) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetUint")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%d", key, value)
	return set(key, fmt.Sprintf("%d", value))
}

func SetUint32(key string, value uint32) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetUint32")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%d", key, value)
	return set(key, fmt.Sprintf("%d", value))
}

func SetUint64(key string, value uint64) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetUint64")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%d", key, value)
	return set(key, fmt.Sprintf("%d", value))
}

func SetIntSlice(key string, value []int) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetIntSlice")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%v", key, value)
	valuesText := []string{}
	var s string
	for _, i := range value {
		s = fmt.Sprintf("%d", i)
		valuesText = append(valuesText, s)
	}
	return set(key, strings.Join(valuesText, "+"))
}

func SetString(key string, value string) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetString")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%s", key, value)
	return set(key, value)
}

func SetStringMap(key string, value map[string]interface{}) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetStringMap")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%v", key, value)
	stringValueJson, _ := json.Marshal(value)
	return set(key, string(stringValueJson))
}

func SetStringMapString(key string, value map[string]string) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetStringMapString")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%v", key, value)
	stringValueJson, _ := json.Marshal(value)
	return set(key, string(stringValueJson))
}

func SetStringSlice(key string, value []string) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetStringSlice")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%v", key, value)
	return set(key, strings.Join(value, ","))
}

func SetTime(key string, value time.Time) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.SetTime")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%s", key, value.Format("2006-01-02T15:04"))
	return set(key, strconv.FormatInt(value.Unix(), 10))
}

func SetDuration(key string, value time.Duration) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.Set")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s, value=%s", key, helpers.DurationToString(value))
	return set(key, value.String())
}
