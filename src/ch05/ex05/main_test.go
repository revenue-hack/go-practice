package main

import (
	"golang.org/x/net/html"
	"net/http"
	"testing"
)

func TestCountWordsAndImages(t *testing.T) {
	resp, err := http.Get("https://www.yahoo.co.jp")
	if err != nil {
		t.Errorf("get HTML: %s", err)
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Errorf("parsing HTML: %s", err)
		return
	}
	words, images := countWordsAndImages(doc)
	if words != 545 || images != 55 {
		t.Errorf("countWordsANdImages error %v\t%v\n", words, images)
	}
}
