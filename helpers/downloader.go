package helpers

import (
	"io"
	"net/http"
	"os"
)

type WriteCounter struct {
	FileSize       uint64
	Total          uint64
	ReportProgress func(wc *WriteCounter)
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.ReportProgress(wc)
	return n, nil
}

var DownloadFile = downloadFile

func downloadFile(filepath string, url string, counter *WriteCounter) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		defer os.Remove(filepath + ".tmp")
		return err
	}
	defer resp.Body.Close()

	io.Copy(out, io.TeeReader(resp.Body, counter))
	os.Rename(filepath+".tmp", filepath)
	os.Chmod(filepath, 0700)

	return nil
}

type UpdateProgress struct {
	Step  int
	Total int
	Info  string
}
