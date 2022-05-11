package file_test

import (
	"os"
	"testing"

	"github.com/Tackem-org/Global/file"
	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	os.Create("test.file")
	assert.True(t, file.FileExists("test.file"))
	assert.False(t, file.FileExists("missing.test.file"))
	os.Remove("test.file")
}
