package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func main() {

	words, images, err := CountWordsAndImages("https://www.yahoo.co.jp")
	if err != nil {
		panic(err)
	}
	fmt.Printf("words: %d\timages: %d\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(doc *html.Node) (words, images int) {
	if doc != nil {
		in := bufio.NewScanner(strings.NewReader(doc.Data))
		in.Split(bufio.ScanWords)
		for in.Scan() {
			words++
		}
		if doc.Type == html.ElementNode && doc.Data == "img" {
			images++
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		word, image := countWordsAndImages(c)
		words += word
		images += image
	}
	return
}
