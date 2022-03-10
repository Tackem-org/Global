package config

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	configClient "github.com/Tackem-org/Global/system/grpcSystem/clients/config"
	pb "github.com/Tackem-org/Proto/pb/config"

	str2duration "github.com/xhit/go-str2duration/v2"
)

func GetBool(key string) (bool, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetBool")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return false, err
	}
	val, err := strconv.ParseBool(response.GetValue())
	if err != nil {
		return false, err
	}
	return val, nil
}

func GetFloat64(key string) (float64, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetFloat64")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0.0, err
	}
	val, err := strconv.ParseFloat(response.GetValue(), 64)
	if err != nil {
		return 0.0, err
	}
	return val, nil
}

func GetInt(key string) (int, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetInt")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0, err
	}
	val, err := strconv.ParseInt(response.GetValue(), 10, 64)
	if err != nil {
		return 0, err
	}
	return int(val), nil
}

func GetInt32(key string) (int32, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetInt32")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0, err
	}
	val, err := strconv.ParseInt(response.GetValue(), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

func GetInt64(key string) (int64, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetInt64")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0, err
	}
	val, err := strconv.ParseInt(response.GetValue(), 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(val), nil
}

func GetUint(key string) (uint, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetUint")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0, err
	}
	val, err := strconv.ParseUint(response.GetValue(), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}

func GetUint32(key string) (uint32, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetUint32")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0, err
	}
	val, err := strconv.ParseUint(response.GetValue(), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

func GetUint64(key string) (uint64, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetUint64")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0, err
	}
	val, err := strconv.ParseUint(response.GetValue(), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(val), nil
}

func GetIntSlice(key string) ([]int, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetIntSlice")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return []int{}, err
	}
	r := []int{}
	for _, n := range strings.Split(response.Value, "+") {
		i, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			return []int{}, err
		}
		r = append(r, int(i))
	}
	return r, nil
}

func GetString(key string) (string, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetString")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return "", err
	}
	return response.GetValue(), nil
}

func GetStringMap(key string) (map[string]interface{}, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetStringMap")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	r := map[string]interface{}{}
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return r, err
	}
	err = json.Unmarshal([]byte(response.GetValue()), &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func GetStringMapString(key string) (map[string]string, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetStringMapString")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	r := map[string]string{}
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return r, err
	}
	err = json.Unmarshal([]byte(response.GetValue()), &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func GetStringSlice(key string) ([]string, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetStringSlice")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return []string{}, err
	}
	return strings.Split(response.Value, ","), nil
}

func GetTime(key string) (time.Time, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetTime")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return time.Now(), err
	}

	val, err := strconv.ParseInt(response.GetValue(), 10, 64)
	if err != nil {
		return time.Now(), err
	}
	return time.Unix(val, 0), nil
}

func GetDuration(key string) (time.Duration, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.config.GetDuration")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] key=%s", key)
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return time.Duration(0), err
	}
	duration, err := str2duration.ParseDuration(response.Value)
	if err != nil {
		return time.Duration(0), err
	}
	return duration, nil
}
