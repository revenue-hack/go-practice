package main

import (
	"net/url"
	"testing"
)

func TestCreateDir(t *testing.T) {
	strings := map[string]bool{
		"hoge": true,
		"":     false,
	}
	for path, is := range strings {
		if createDir(path) != is {
			t.Errorf("%s\t%v\n", path, is)
		}
	}
}

func TestIsTarget(t *testing.T) {
	maps := map[string]bool{
		"https://golang.org": true,
		"":                   false,
		"https://golang.com": false,
	}
	domain, err := url.Parse("https://golang.org")
	if err != nil {
		t.Errorf("%v\n", domain)
		return
	}
	for url, is := range maps {
		if isTarget(url, domain) != is {
			t.Errorf("url: %s\tdomain: %v\n", url, domain)
			return
		}
	}
}

func TestFileName(t *testing.T) {
	maps := map[string]string{
		"hoge": "hoge",
		"":     "",
		".":    "index",
		"/":    "index",
	}
	for name, expected := range maps {
		if fileName(name) != expected {
			t.Errorf("name: %s\texpected: %s\n", name, expected)
		}
	}
}

func TestDir(t *testing.T) {
	maps := map[string]string{
		"hoge": "hoge",
		"":     "",
		".":    "",
		"/":    "",
	}
	for name, expected := range maps {
		if dir(name) != expected {
			t.Errorf("name: %s\texpected: %s\n", name, expected)
		}
	}
}
