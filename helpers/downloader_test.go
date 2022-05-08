package helpers_test

import (
	"os"
	"testing"

	"github.com/Tackem-org/Global/helpers"
	"github.com/stretchr/testify/assert"
)

func TestWriteCounter_Write(t *testing.T) {
	counter := &helpers.WriteCounter{
		FileSize:       200,
		Total:          0,
		ReportProgress: func(wc *helpers.WriteCounter) {},
	}

	l, err := counter.Write([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	assert.Nil(t, err)
	assert.Equal(t, uint64(10), counter.Total)
	assert.Equal(t, 10, l)
}

func TestDownloadFile(t *testing.T) {
	counter := &helpers.WriteCounter{
		FileSize:       0,
		Total:          0,
		ReportProgress: func(wc *helpers.WriteCounter) {},
	}
	errPass := helpers.DownloadFile("test.txt", "https://raw.githubusercontent.com/Tackem-org/Repo/master/LICENSE.md", counter)
	assert.Nil(t, errPass)
	assert.Equal(t, uint64(15174), counter.Total)
	os.Remove("test.txt")

	err1 := helpers.DownloadFile("/fail", "https://raw.githubusercontent.com/Tackem-org/Repo/master/LICENSE.md", counter)
	assert.Error(t, err1)

	err2 := helpers.DownloadFile("test.txt", "", counter)
	assert.Error(t, err2)
}
