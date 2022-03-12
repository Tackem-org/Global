package healthCheck_test

import (
	"testing"

	"github.com/Tackem-org/Global/healthCheck"
	"github.com/stretchr/testify/assert"
)

func TestIssues(t *testing.T) {
	assert.Empty(t, healthCheck.Issues())
	healthCheck.SetIssue("TEST")
	assert.NotEmpty(t, healthCheck.Issues())
	healthCheck.ClearIssue("TEST")
	assert.Empty(t, healthCheck.Issues())
}
func TestHealthy(t *testing.T) {
	assert.True(t, healthCheck.Healthy())
	healthCheck.SetIssue("TEST")
	assert.False(t, healthCheck.Healthy())
}
