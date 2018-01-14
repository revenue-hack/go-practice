package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("args len is not 2")
		os.Exit(1)
	}
	resp, err := http.Get(os.Args[1])
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink1: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", ElementsByTagName(doc, "img", "h2", "h1", "div"))
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var elements []*html.Node
	if doc.Type == html.ElementNode {
		for _, n := range name {
			if doc.Data == n {
				elements = append(elements, doc)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		for _, element := range ElementsByTagName(c, name...) {
			elements = append(elements, element)
		}
	}
	return elements
}
