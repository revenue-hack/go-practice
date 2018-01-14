package main

import (
	"path"
	"os"
	"io"
	"fmt"
	"net/http"
)

func main() {
	fileName, n, err := fetch("https://golang.org")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\t%v\n", fileName, n)

}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}

