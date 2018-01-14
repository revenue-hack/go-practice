package main

import (
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
	if n := ElementByID(doc, os.Args[2]); n != nil {
		fmt.Printf("find: %s\n", n.Data)
	} else {
		fmt.Println("nothing")
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	n, has := forEachNode(doc, id, findById, nil)
	if has {
		return n
	} else {
		return nil
	}
}

func findById(doc *html.Node, id string) bool {
	if doc.Type == html.ElementNode {
		attrs := doc.Attr
		for _, attr := range attrs {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
	}
	return false
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) (*html.Node, bool) {
	if pre != nil && pre(n, id) {
		return n, true
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if rn, isTrue := forEachNode(c, id, pre, post); isTrue {
			return rn, true
		}
	}
	if post != nil {
		post(n, id)
	}
	return nil, false
}
