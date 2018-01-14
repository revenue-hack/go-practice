package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
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
	forEachNode(doc, start, end)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func start(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		writeStartElementNode(n)
	case html.TextNode:
		writeStartTextNode(n)
	case html.DocumentNode:
		fmt.Println("<!DOCTYPE html>")
	}
}

func writeStartElementNode(n *html.Node) {
	depth++
	attr := createAttr(n.Attr)
	if attr != "" {
		if n.Data == "img" {
			fmt.Printf("\n%*s<%s %s/>\n", depth*2, "", n.Data, attr)
			depth--
		} else {
			fmt.Printf("\n%*s<%s %s>", depth*2, "", n.Data, attr)
		}
	} else if n.Data == "br" {
		fmt.Printf("\n%*s<%s />", depth*2, "", n.Data)
		depth--
	} else {
		fmt.Printf("\n%*s<%s>", depth*2, "", n.Data)
	}
}

func writeStartTextNode(n *html.Node) {
	fmt.Printf("%s", n.Data)
}

func createAttr(attrs []html.Attribute) string {
	var buf bytes.Buffer
	if attrs == nil {
		return buf.String()
	}
	for i, attr := range attrs {
		buf.WriteString(attr.Key)
		buf.WriteString(`="`)
		buf.WriteString(attr.Val)
		buf.WriteString(`"`)
		if i != 0 {
			buf.WriteString(" ")
		}
	}
	return buf.String()
}

func end(n *html.Node) {
	if n.Type == html.ElementNode && n.Data != "br" && n.Data != "img" {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		depth--
	}
}
