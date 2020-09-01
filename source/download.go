package source

import (
	"io"
	"net/http"
)

//DownloadURL downloads from the specified url
func DownloadURL(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
