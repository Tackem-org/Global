package sysErrors

import (
	"fmt"
)

type MasterDownError struct{ err string }

func (e *MasterDownError) Error() string {
	return fmt.Sprintf("master is down %s", e.err)
}

type ServiceDownError struct{ err string }

func (e *ServiceDownError) Error() string {
	return fmt.Sprintf("service is down %s", e.err)
}

type ServiceInactiveError struct{ err string }

func (e *ServiceInactiveError) Error() string {
	return fmt.Sprintf("service is inactive %s", e.err)
}

type ConfigTypeError struct{ err string }

func (e *ConfigTypeError) Error() string {
	return fmt.Sprintf("config value type is wrong %s", e.err)
}

type ConfigValueError struct{ err string }

func (e *ConfigValueError) Error() string {
	return fmt.Sprintf("config value is bad %s", e.err)
}
