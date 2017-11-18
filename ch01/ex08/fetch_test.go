package main

import (
	"testing"
)

func TestPrefixHttpForHttp(t *testing.T) {
	if prefixHttp("http://www.hoge.com") != "http://www.hoge.com" {
		t.Error("TestPrefixHttpForHttp func error")
	}
}

func TestPrefixHttpForNonHttp(t *testing.T) {
	if prefixHttp("www.hoge.com") != "http://www.hoge.com" {
		t.Error("TestPrefixHttpForHttp func error")
	}
}
