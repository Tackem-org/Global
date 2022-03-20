package config

import (
	"strconv"
	"strings"
	"time"

	"github.com/Tackem-org/Global/sysErrors"
	configClient "github.com/Tackem-org/Global/system/grpcSystem/clients/config"
	pb "github.com/Tackem-org/Proto/pb/config"

	str2duration "github.com/xhit/go-str2duration/v2"
)

//TODO FINISH THIS OFF WITH BAD DATA COMING IN TO GET ERROR DATA
func GetBool(key string) (bool, error) {
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return false, err
	}
	if response.Type != pb.ValueType_Bool {
		return false, &sysErrors.ConfigTypeError{}
	}
	val, err := strconv.ParseBool(response.GetValue())
	if err != nil {
		return false, &sysErrors.ConfigValueError{}
	}
	return val, nil
}

func GetFloat64(key string) (float64, error) {
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0.0, err
	}
	if response.Type != pb.ValueType_Float64 {
		return 0.0, &sysErrors.ConfigTypeError{}
	}
	val, err := strconv.ParseFloat(response.GetValue(), 64)
	if err != nil {
		return 0.0, &sysErrors.ConfigValueError{}
	}
	return val, nil
}

func GetInt(key string) (int, error) {
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0, err
	}
	if response.Type != pb.ValueType_Int {
		return 0, &sysErrors.ConfigTypeError{}
	}
	val, err := strconv.ParseInt(response.GetValue(), 10, 64)
	if err != nil {
		return 0, &sysErrors.ConfigValueError{}
	}
	return int(val), nil
}

func GetInt32(key string) (int32, error) {
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0, err
	}
	if response.Type != pb.ValueType_Int32 {
		return 0, &sysErrors.ConfigTypeError{}
	}
	val, err := strconv.ParseInt(response.GetValue(), 10, 32)
	if err != nil {
		return 0, &sysErrors.ConfigValueError{}
	}
	return int32(val), nil
}

func GetInt64(key string) (int64, error) {
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0, err
	}
	if response.Type != pb.ValueType_Int64 {
		return 0, &sysErrors.ConfigTypeError{}
	}
	val, err := strconv.ParseInt(response.GetValue(), 10, 64)
	if err != nil {
		return 0, &sysErrors.ConfigValueError{}
	}
	return int64(val), nil
}

func GetUint(key string) (uint, error) {
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0, err
	}
	if response.Type != pb.ValueType_Uint {
		return 0, &sysErrors.ConfigTypeError{}
	}
	val, err := strconv.ParseUint(response.GetValue(), 10, 64)
	if err != nil {
		return 0, &sysErrors.ConfigValueError{}
	}
	return uint(val), nil
}

func GetUint32(key string) (uint32, error) {
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0, err
	}
	if response.Type != pb.ValueType_Uint32 {
		return 0, &sysErrors.ConfigTypeError{}
	}
	val, err := strconv.ParseUint(response.GetValue(), 10, 32)
	if err != nil {
		return 0, &sysErrors.ConfigValueError{}
	}
	return uint32(val), nil
}

func GetUint64(key string) (uint64, error) {
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return 0, err
	}
	if response.Type != pb.ValueType_Uint64 {
		return 0, &sysErrors.ConfigTypeError{}
	}
	val, err := strconv.ParseUint(response.GetValue(), 10, 64)
	if err != nil {
		return 0, &sysErrors.ConfigValueError{}
	}
	return uint64(val), nil
}

func GetIntSlice(key string) ([]int, error) {
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return []int{}, err
	}
	if response.Type != pb.ValueType_IntSlice {
		return []int{}, &sysErrors.ConfigTypeError{}
	}
	r := []int{}
	for _, n := range strings.Split(response.Value, "+") {
		i, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			return []int{}, &sysErrors.ConfigValueError{}
		}
		r = append(r, int(i))
	}
	return r, nil
}

func GetString(key string) (string, error) {
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return "", err
	}
	if response.Type != pb.ValueType_String {
		return "", &sysErrors.ConfigTypeError{}
	}
	return response.GetValue(), nil
}

func GetStringSlice(key string) ([]string, error) {
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return []string{}, err
	}
	if response.Type != pb.ValueType_StringSlice {
		return []string{}, &sysErrors.ConfigTypeError{}
	}
	return strings.Split(response.Value, ","), nil
}

func GetTime(key string) (time.Time, error) {
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return time.Unix(0, 0), err
	}
	if response.Type != pb.ValueType_Time {
		return time.Unix(0, 0), &sysErrors.ConfigTypeError{}
	}
	val, err := strconv.ParseInt(response.GetValue(), 10, 64)
	if err != nil {
		return time.Unix(0, 0), &sysErrors.ConfigValueError{}
	}
	return time.Unix(val, 0), nil
}

func GetDuration(key string) (time.Duration, error) {
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return time.Duration(0), err
	}
	if response.Type != pb.ValueType_Duration {
		return time.Duration(0), &sysErrors.ConfigTypeError{}
	}
	duration, err := str2duration.ParseDuration(response.Value)
	if err != nil {
		return time.Duration(0), &sysErrors.ConfigValueError{}
	}
	return duration, nil
}
