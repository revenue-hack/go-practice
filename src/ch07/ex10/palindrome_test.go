package main

import "testing"

var data = []struct{
	s string
	expected bool
}{
	{"hhhh", true},
	{"", true},
	{"heello", false},
	{"aaaaaa", true},
	{"abccba", true},
	{"1122", false},
}
func TestIsPalindrome(t *testing.T) {
	for _, d := range data {
		b := []byte(d.s)
		if IsPalindrome(Byte(b)) != d.expected {
			t.Errorf("don't match\tstring: %s\texpected: %v\n", d.s, d.expected)
		}
	}
}

