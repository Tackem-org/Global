package system

import "fmt"

type SystemDownError struct{ err string }

func (e *SystemDownError) Error() string {
	return fmt.Sprintf("System Down Unable to Login %s", e.err)
}
