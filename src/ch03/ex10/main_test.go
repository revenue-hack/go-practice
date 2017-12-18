package main

import "testing"

func TestCommaOfNormal(t *testing.T) {
	s := "aaabbbccc"
	if comma(s) != "aaa,bbb,ccc" {
		t.Error("Test Normal Error")
	}
}

func TestCommaOfRemainderNormal(t *testing.T) {
	s := "aabbbccc"
	if comma(s) != "aa,bbb,ccc" {
		t.Error("Test RmainderNormal Error")
	}
}

func TestCommaOfEmpty(t *testing.T) {
	s := ""
	if comma(s) != "" {
		t.Error("Test Empty Error")
	}
}
