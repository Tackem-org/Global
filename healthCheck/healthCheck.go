package healthCheck

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

var (
	healthy bool = true
	issues  []string
)

func Healthy() bool {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.healthCheck.Healthy")
	return healthy
}

func Issues() []string {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.healthCheck.Issues")
	return issues
}

func SetIssue(issue string) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.healthCheck.SetIssues")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] issue=%s", issue)
	issues = append(issues, issue)
}
