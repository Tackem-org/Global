package system

import "fmt"

type ServiceDownError struct{ err string }

func (e *ServiceDownError) Error() string {
	return fmt.Sprintf("Service Down Unable to Login %s", e.err)
}
