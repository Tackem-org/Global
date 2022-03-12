package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	configClient "github.com/Tackem-org/Global/system/grpcSystem/clients/config"
	pb "github.com/Tackem-org/Proto/pb/config"
)

func set(key string, value string) (bool, error) {
	response, _ := configClient.Set(&pb.SetConfigRequest{Key: key, Value: value})
	return response.GetSuccess(), errors.New(response.GetErrorMessage())
}

func SetBool(key string, value bool) (bool, error) {
	return set(key, fmt.Sprintf("%t", value))
}

func SetFloat64(key string, value float64) (bool, error) {
	return set(key, fmt.Sprintf("%f", value))
}

func SetInt(key string, value int) (bool, error) {
	return set(key, fmt.Sprintf("%d", value))
}

func SetInt32(key string, value int32) (bool, error) {
	return set(key, fmt.Sprintf("%d", value))
}

func SetInt64(key string, value int64) (bool, error) {
	return set(key, fmt.Sprintf("%d", value))
}

func SetUint(key string, value uint) (bool, error) {
	return set(key, fmt.Sprintf("%d", value))
}

func SetUint32(key string, value uint32) (bool, error) {
	return set(key, fmt.Sprintf("%d", value))
}

func SetUint64(key string, value uint64) (bool, error) {
	return set(key, fmt.Sprintf("%d", value))
}

func SetIntSlice(key string, value []int) (bool, error) {
	valuesText := []string{}
	var s string
	for _, i := range value {
		s = fmt.Sprintf("%d", i)
		valuesText = append(valuesText, s)
	}
	return set(key, strings.Join(valuesText, "+"))
}

func SetString(key string, value string) (bool, error) {
	return set(key, value)
}

func SetStringMap(key string, value map[string]interface{}) (bool, error) {
	stringValueJson, _ := json.Marshal(value)
	return set(key, string(stringValueJson))
}

func SetStringMapString(key string, value map[string]string) (bool, error) {
	stringValueJson, _ := json.Marshal(value)
	return set(key, string(stringValueJson))
}

func SetStringSlice(key string, value []string) (bool, error) {
	return set(key, strings.Join(value, ","))
}

func SetTime(key string, value time.Time) (bool, error) {
	return set(key, strconv.FormatInt(value.Unix(), 10))
}

func SetDuration(key string, value time.Duration) (bool, error) {
	return set(key, value.String())
}
