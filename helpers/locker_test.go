package helpers_test

import (
	"testing"
	"time"

	"github.com/Tackem-org/Global/helpers"
	"github.com/stretchr/testify/assert"
)

func TestLocker(t *testing.T) {
	l := helpers.Locker{Label: "TEST"}
	l.Down()
	assert.False(t, l.Check())
	l.Up()
	assert.True(t, l.Check())
	l.Down()
	assert.False(t, l.Wait(time.Duration(time.Microsecond)))
	l.Up()
	assert.True(t, l.Wait(time.Duration(time.Microsecond)))
}
