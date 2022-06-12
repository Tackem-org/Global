package channels_test

import (
	"testing"

	"github.com/Tackem-org/Global/system/channels"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	channels.Setup()
	assert.NotNil(t, channels.Root.Shutdown)
	assert.NotNil(t, channels.Root.TermChan)
}

func TestCallShutdown(t *testing.T) {
	channels.Setup()
	channels.CallShutdown()
	assert.NotNil(t, channels.Root.Shutdown)

}
