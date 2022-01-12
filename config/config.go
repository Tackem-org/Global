package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/system"
	pb "github.com/Tackem-org/Proto/pb/config"

	str2duration "github.com/xhit/go-str2duration/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func get(key string) (*pb.GetConfigResponse, error) {
	conn, err := system.GetMasterConnection(false)
	if err != nil {
		logging.Error("[Remote Config Get] Cannot Connect to the Server: " + err.Error())
		return nil, err
	}
	defer conn.Close()

	client := pb.NewConfigClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, metadata.MD{})

	response, err := client.Get(ctx, &pb.GetConfigRequest{Key: key}, grpc.Header(&metadata.MD{}))
	if err != nil {
		logging.Error("[Remote Config Get] Error with the Servers Get: " + err.Error())
		return nil, err
	}
	return response, nil
}

func GetBool(key string) (bool, error) {
	response, err := get(key)
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
	response, err := get(key)
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
	response, err := get(key)
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
	response, err := get(key)
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
	response, err := get(key)
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
	response, err := get(key)
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
	response, err := get(key)
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
	response, err := get(key)
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
	response, err := get(key)
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
	response, err := get(key)
	if err != nil {
		return "", err
	}
	return response.GetValue(), nil
}

func GetStringMap(key string) (map[string]interface{}, error) {
	r := map[string]interface{}{}
	response, err := get(key)
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
	r := map[string]string{}
	response, err := get(key)
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
	response, err := get(key)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(response.Value, ","), nil
}

func GetTime(key string) (time.Time, error) {
	response, err := get(key)
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
	response, err := get(key)
	if err != nil {
		return time.Duration(0), err
	}
	duration, err := str2duration.ParseDuration(response.Value)
	if err != nil {
		return time.Duration(0), err
	}
	return duration, nil
}

func set(key string, value string) (bool, error) {
	conn, err := system.GetMasterConnection(false)

	if err != nil {
		logging.Error("[Remote Config Set] Cannot Connect to the Server: " + err.Error())
		return false, err
	}
	defer conn.Close()

	client := pb.NewConfigClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, metadata.MD{})

	response, err := client.Set(ctx, &pb.SetConfigRequest{Key: key, Value: value}, grpc.Header(&metadata.MD{}))
	if err != nil {
		logging.Error("[Remote Config Set] Error with the Servers Set: " + err.Error())
		return false, err
	}
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
