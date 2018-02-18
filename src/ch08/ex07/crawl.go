package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"gopl.io/ch5/links"
)

func main() {
	domain, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Errorf("%v\n", domain)
		panic(err)
	}
	worklist := make(chan []string)
	go func() { worklist <- []string{os.Args[1]} }()
	var n int
	n++
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				if isTarget(link, domain) {
					go func(link string) {
						download(link)
						worklist <- crawl(link)
					}(link)
				}
			}
		}
	}
}

func createDir(path string) bool {
	if err := os.MkdirAll(path, 0777); err != nil {
		return false
	}
	return true
}

func isTarget(item string, domain *url.URL) bool {
	parseUrl, err := url.Parse(item)
	if err != nil {
		panic(err)
	}
	return strings.HasSuffix(parseUrl.Host, domain.Host)
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func fileName(fileName string) string {
	if fileName == "." || fileName == "/" {
		fileName = "index"
	}
	return fileName
}

func dir(dir string) string {
	if dir == "." || dir == "/" {
		dir = ""
	}
	return dir
}

func download(url string) bool {
	fmt.Printf("download: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dir := dir(path.Dir(resp.Request.URL.Path))
	host := resp.Request.Host
	dirPath := fmt.Sprintf("%s%s", host, dir)
	if !createDir(dirPath) {
		fmt.Errorf("%s\n", "func createDomainDirs error")
	}
	fileName := fileName(path.Base(url))
	filePath := fmt.Sprintf("%s/%s", dirPath, fileName)
	if file, err := os.Create(filePath); err == nil {
		defer file.Close()
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("%s donwlaod continue\n", url)
	}
	return true
}
