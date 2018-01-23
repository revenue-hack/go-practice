package main

import (
	"golang.org/x/net/html"
	"testing"
)

func TestVisit(t *testing.T) {
	contents := `<html><a href="https://yahoo.co.jp"></a><a href="https://golang.org"></a></html>`
	doc, err := html.Parse(newReader(contents))
	if err != nil {
		panic(err)
	}
	links := []string{"https://yahoo.co.jp", "https://golang.org"}
	for i, link := range visit(nil, doc) {
		if links[i] != link {
			t.Errorf("visit link don't match visit link: %s\t expected link: %s\n", link, links[i])
		}
	}
}
