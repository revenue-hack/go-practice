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
	result := make(map[string]int)
	if len(visit(result, doc)) != 51 {
		t.Errorf("visit length: %d\n", len(visit(result, doc)))
	}
}
