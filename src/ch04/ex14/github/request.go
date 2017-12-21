package github

import (
	"io"
	"net/http"
)

func createRequest(method, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	return request
}

func get(url string) (*http.Response, error) {
	request := createRequest("GET", url, nil)
	return http.DefaultClient.Do(request)
}
