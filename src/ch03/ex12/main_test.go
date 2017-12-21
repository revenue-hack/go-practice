package main

import "testing"

func TestAnagramOfNormal(t *testing.T) {
	s1 := "aaabbb"
	s2 := "bbaaba"
	if !anagram(s1, s2) {
		t.Error("Test Normal Error")
	}

	s1 = "aaa"
	s2 = "bbb"
	if anagram(s1, s2) {
		t.Error("Test Normal Error 2")
	}
}

func TestAnagramOfOneEmpty(t *testing.T) {
	s1 := ""
	s2 := "ssjojodw"
	if anagram(s1, s2) {
		t.Error("Test One Empty Error")
	}
	if anagram(s2, s1) {
		t.Error("Test One Empty Error")
	}
}
