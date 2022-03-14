package config

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	configClient "github.com/Tackem-org/Global/system/grpcSystem/clients/config"
	pb "github.com/Tackem-org/Proto/pb/config"

	str2duration "github.com/xhit/go-str2duration/v2"
)

func GetBool(key string) (bool, error) {
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
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return "", err
	}
	return response.GetValue(), nil
}

//TODO CHANGE WITH NEW FUNC BELLOW AFTER TESTED
func GetStringMap(key string) (map[string]interface{}, error) {
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

//NEW
// func GetStringMap(key string) (map[string]interface{}, error) {

// 	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
// 	if err != nil {
// 		return map[string]interface{}{}, err
// 	}
// 	r, err := helpers.StringToStringMap(response.Value)
// 	if err != nil {
// 		return map[string]interface{}{}, err
// 	}
// 	return r, nil
// }

func GetStringMapString(key string) (map[string]string, error) {
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
	response, err := configClient.Get(&pb.GetConfigRequest{Key: key})
	if err != nil {
		return []string{}, err
	}
	return strings.Split(response.Value, ","), nil
}

func GetTime(key string) (time.Time, error) {
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
