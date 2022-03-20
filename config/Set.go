package config

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	configClient "github.com/Tackem-org/Global/system/grpcSystem/clients/config"
	pb "github.com/Tackem-org/Proto/pb/config"
)

func SetBool(key string, value bool) (bool, error) {
	response, err := configClient.Set(&pb.SetConfigRequest{Key: key, Value: fmt.Sprintf("%t", value)})
	if !response.Success {
		err = errors.New(response.ErrorMessage)
	}
	return response.Success, err
}

func SetFloat64(key string, value float64) (bool, error) {
	response, err := configClient.Set(&pb.SetConfigRequest{Key: key, Value: fmt.Sprintf("%f", value)})
	if !response.Success {
		err = errors.New(response.ErrorMessage)
	}
	return response.Success, err
}

func SetInt(key string, value int) (bool, error) {
	response, err := configClient.Set(&pb.SetConfigRequest{Key: key, Value: fmt.Sprintf("%d", value)})
	if !response.Success {
		err = errors.New(response.ErrorMessage)
	}
	return response.Success, err
}

func SetInt32(key string, value int32) (bool, error) {
	response, err := configClient.Set(&pb.SetConfigRequest{Key: key, Value: fmt.Sprintf("%d", value)})
	if !response.Success {
		err = errors.New(response.ErrorMessage)
	}
	return response.Success, err
}

func SetInt64(key string, value int64) (bool, error) {
	response, err := configClient.Set(&pb.SetConfigRequest{Key: key, Value: fmt.Sprintf("%d", value)})
	if !response.Success {
		err = errors.New(response.ErrorMessage)
	}
	return response.Success, err
}

func SetUint(key string, value uint) (bool, error) {
	response, err := configClient.Set(&pb.SetConfigRequest{Key: key, Value: fmt.Sprintf("%d", value)})
	if !response.Success {
		err = errors.New(response.ErrorMessage)
	}
	return response.Success, err
}

func SetUint32(key string, value uint32) (bool, error) {
	response, err := configClient.Set(&pb.SetConfigRequest{Key: key, Value: fmt.Sprintf("%d", value)})
	if !response.Success {
		err = errors.New(response.ErrorMessage)
	}
	return response.Success, err
}

func SetUint64(key string, value uint64) (bool, error) {
	response, err := configClient.Set(&pb.SetConfigRequest{Key: key, Value: fmt.Sprintf("%d", value)})
	if !response.Success {
		err = errors.New(response.ErrorMessage)
	}
	return response.Success, err
}

func SetIntSlice(key string, value []int) (bool, error) {
	valuesText := []string{}
	var s string
	for _, i := range value {
		s = fmt.Sprintf("%d", i)
		valuesText = append(valuesText, s)
	}
	response, err := configClient.Set(&pb.SetConfigRequest{Key: key, Value: strings.Join(valuesText, "+")})
	if !response.Success {
		err = errors.New(response.ErrorMessage)
	}
	return response.Success, err
}

func SetString(key string, value string) (bool, error) {
	response, err := configClient.Set(&pb.SetConfigRequest{Key: key, Value: value})
	if !response.Success {
		err = errors.New(response.ErrorMessage)
	}
	return response.Success, err
}

func SetStringSlice(key string, value []string) (bool, error) {
	response, err := configClient.Set(&pb.SetConfigRequest{Key: key, Value: strings.Join(value, ",")})
	if !response.Success {
		err = errors.New(response.ErrorMessage)
	}
	return response.Success, err
}

func SetTime(key string, value time.Time) (bool, error) {
	response, err := configClient.Set(&pb.SetConfigRequest{Key: key, Value: strconv.FormatInt(value.Unix(), 10)})
	if !response.Success {
		err = errors.New(response.ErrorMessage)
	}
	return response.Success, err
}

func SetDuration(key string, value time.Duration) (bool, error) {
	response, err := configClient.Set(&pb.SetConfigRequest{Key: key, Value: value.String()})
	if !response.Success {
		err = errors.New(response.ErrorMessage)
	}
	return response.Success, err
}
