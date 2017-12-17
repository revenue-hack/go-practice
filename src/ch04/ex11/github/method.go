package github

import (
	"io"
	"net/http"
)

func post(url string, body io.Reader) (*http.Response, error) {
	request := createRequest("POST", url, body)
	return http.DefaultClient.Do(request)
}

func get(url string) (*http.Response, error) {
	request := createRequest("GET", url, nil)
	return http.DefaultClient.Do(request)
}

func patch(url string, body io.Reader) (*http.Response, error) {
	request := createRequest("PATCH", url, body)
	return http.DefaultClient.Do(request)
}

func createRequest(method, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	credential := envLoad()
	request.SetBasicAuth(credential.user, credential.password)
	request.Header.Set("Accept", contentType)
	return request
}
