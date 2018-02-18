package main

import (
	"fmt"
	"net/http"
	"os"
)

var done = make(chan struct{})

func main() {
	lists := make(chan []string)
	go func() { lists <- os.Args[1:] }()
	response := make(chan string)
	for _, url := range <-lists {
		go fetch(url, response)
	}
	fmt.Printf("first url: %s\n", <-response)
	close(done)
}
func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
func fetch(url string, responseCh chan string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("%v\n", err)
		return
	}
	if canceled() {
		return
	}
	go func() {
		select {
		case <-done:
			req.Close = true
		}
	}()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("%v\n", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Errorf("getting %s: %s", url, resp.Status)
		return
	}
	select {
	case <-done:
	case responseCh <- url:
	}
}
