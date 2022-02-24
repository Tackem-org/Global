package sysErrors

import "fmt"

type ServiceDownError struct{ err string }

func (e *ServiceDownError) Error() string {
	return fmt.Sprintf("service is down %s", e.err)
}

type ServiceInactiveError struct{ err string }

func (e *ServiceInactiveError) Error() string {
	return fmt.Sprintf("service is inactive %s", e.err)
}
