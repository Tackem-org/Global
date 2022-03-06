package healthCheck

import (
	"sync"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

var (
	mu      sync.RWMutex
	healthy bool = true
	issues  []string
)

func Healthy() bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.healthCheck.Healthy")
	mu.RLock()
	defer mu.RUnlock()
	return healthy
}

func Issues() []string {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.healthCheck.Issues")
	mu.RLock()
	defer mu.RUnlock()
	return issues
}

func SetIssue(issue string) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.healthCheck.SetIssues")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] issue=%s", issue)
	mu.Lock()
	defer mu.Unlock()
	issues = append(issues, issue)
}
