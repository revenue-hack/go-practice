package main

import (
	"golang.org/x/net/html"
	"os"
	"testing"
)

func TestElementsByTagName(t *testing.T) {
	file, err := os.Open("index.html")
	if err != nil {
		t.Errorf("file open error\n%v\n", err)
	}
	defer file.Close()
	doc, err := html.Parse(file)
	if err != nil {
		t.Errorf("file parse error\n%v\n", err)
	}
	h1Nodes := ElementsByTagName(doc, "h1")
	h2Nodes := ElementsByTagName(doc, "h2")
	divNodes := ElementsByTagName(doc, "div")
	imgNodes := ElementsByTagName(doc, "img")
	allNodes := ElementsByTagName(doc, "h1", "h2", "div", "img")
	if len(h1Nodes) != 1 {
		t.Error("h1 node len invalid")
	}
	if len(h2Nodes) != 7 {
		t.Error("h2 node len invalid")
	}
	if len(divNodes) != 273 {
		t.Error("div node len invalid")
	}
	if len(imgNodes) != 75 {
		t.Error("img node len invalid")
	}
	if len(allNodes) != (len(h1Nodes) + len(h2Nodes) + len(divNodes) + len(imgNodes)) {
		t.Error("all node len invalid")
	}
}
