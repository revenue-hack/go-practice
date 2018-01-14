package main

import (
	"golang.org/x/net/html"
	"os"
	"testing"
)

func TestVisit(t *testing.T) {
	file, err := os.Open("index.html")
	if err != nil {
		t.Errorf("open HTML: %s", err)
		return
	}
	defer file.Close()
	doc, err := html.Parse(file)
	if err != nil {
		t.Errorf("parse error: %v\n", err)
		return
	}
	if len(visit(nil, doc)) != 335 {
		t.Errorf("visit length: %d\n", len(visit(nil, doc)))
	}
}
