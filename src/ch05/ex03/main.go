package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

var elements []string

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		elements = append(elements, n.Data)
	}
	if n.Type == html.TextNode && elements[len(elements)-1] != "style" && elements[len(elements)-1] != "script" {
		links = append(links, strings.TrimSpace(n.Data))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
