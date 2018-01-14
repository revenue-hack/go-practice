package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink1: %v\n", err)
		os.Exit(1)
	}
	result := make(map[string]int)
	for tag, num := range visit(result, doc) {
		fmt.Printf("%s\t%d\n", tag, num)
	}
}

func visit(links map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		links[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(links, c)
	}
	return links
}
