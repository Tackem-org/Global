package helpers_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"testing"
	"time"

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

func stringToFile(s string, filename string) {
	file, _ := os.Create(filename)
	defer file.Close()
	file.WriteString(s)
}

func TestDownloadFile(t *testing.T) {
	stringToFile("test", "testfile.txt")
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testfile.txt")
	})

	srv := &http.Server{Addr: ":9999"}
	go srv.ListenAndServe()
	time.Sleep(time.Millisecond)

	counter := &helpers.WriteCounter{
		FileSize:       0,
		Total:          0,
		ReportProgress: func(wc *helpers.WriteCounter) {},
	}

	errPass := helpers.DownloadFile("test.txt", "http://127.0.0.1:9999/test", counter)
	if errPass != nil {
		fmt.Println(errPass.Error())
	}
	assert.Nil(t, errPass)
	assert.Equal(t, uint64(4), counter.Total)
	os.Remove("test.txt")

	if runtime.GOOS == "windows" {
		err1 := helpers.DownloadFile("c:/windows/fail", "http://127.0.0.1:9999/test", counter)
		assert.Error(t, err1)
	} else {
		err1 := helpers.DownloadFile("/fail", "http://127.0.0.1:9999/test", counter)
		assert.Error(t, err1)
	}

	err2 := helpers.DownloadFile("test.txt", "", counter)
	assert.Error(t, err2)
	srv.Shutdown(context.Background())
	os.Remove("testfile.txt")
}
