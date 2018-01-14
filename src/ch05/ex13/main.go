package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

func main() {
	domain, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Errorf("%v\n", domain)
		panic(err)
	}
	breadthFirst(crawl, []string{os.Args[1]}, domain)
}

func createDir(path string) bool {
	if err := os.MkdirAll(path, 0777); err != nil {
		return false
	}
	return true
}

func breadthFirst(f func(item string) []string, worklist []string, domain *url.URL) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				if isTarget(item, domain) {
					download(item)
					worklist = append(worklist, f(item)...)
				}
			}
		}
	}
}

func isTarget(item string, domain *url.URL) bool {
	parseUrl, err := url.Parse(item)
	if err != nil {
		panic(err)
	}
	return strings.HasSuffix(parseUrl.Host, domain.Host)
}

func crawl(url string) []string {
	list, err := links.Extract(url)
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
	filePath := fmt.Sprintf("%s/%s.html", dirPath, fileName)
	if file, err := os.Create(filePath); err == nil {
		defer file.Close()
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("%s donwlaod error\n", url)
	}
	return true
}
