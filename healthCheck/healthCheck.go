package healthCheck

import (
	"sync"
)

var (
	mu     sync.RWMutex
	issues []string
)

func Healthy() bool {
	mu.RLock()
	defer mu.RUnlock()
	return len(issues) == 0
}

func Issues() []string {
	mu.RLock()
	defer mu.RUnlock()
	return issues
}

func SetIssue(issue string) {
	mu.Lock()
	defer mu.Unlock()
	issues = append(issues, issue)
}

func ClearIssue(issue string) {
	mu.Lock()
	defer mu.Unlock()
	for i, v := range issues {
		if v == issue {
			issues = append(issues[:i], issues[i+1:]...)
		}
	}
}
