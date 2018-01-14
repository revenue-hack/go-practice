package main

import (
	"golang.org/x/net/html"
	"os"
	"testing"
)

func TestElementByID(t *testing.T) {
	file, err := os.Open("index.html")
	if err != nil {
		t.Error("index.html error")
	}
	defer file.Close()
	doc, err := html.Parse(file)
	if err != nil {
		t.Error("parse error")
	}
	if ElementByID(doc, "landing_1") == nil {
		t.Error("id: landing_1")
	}
	if ElementByID(doc, "") != nil {
		t.Error("id: empty")
	}
	if ElementByID(doc, "hoge") != nil {
		t.Error("id: hoge")
	}
}
