package sysErrors

import (
	"errors"
)

var (
	MasterDownError      = errors.New("master is down")
	ServiceDownError     = errors.New("service is down")
	ServiceInactiveError = errors.New("service is inactive")
	ConfigTypeError      = errors.New("config value type is wrong")
	ConfigValueError     = errors.New("config value is bad")
)

func SetupErrors() {
	MasterDownError = errors.New("master is down")
	ServiceDownError = errors.New("service is down")
	ServiceInactiveError = errors.New("service is inactive")
	ConfigTypeError = errors.New("config value type is wrong")
	ConfigValueError = errors.New("config value is bad")
}
