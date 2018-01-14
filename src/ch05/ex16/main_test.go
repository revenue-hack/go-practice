package main

import "testing"

var cases = []struct {
	words    []string
	s        string
	expected string
}{
	{[]string{"hello", "world"}, ",", "hello,world"},
	{[]string{"1", "2", "3"}, " ", "1 2 3"},
	{[]string{"hello", "world"}, "", "helloworld"},
	{nil, "", ""},
}

func TestJoin(t *testing.T) {
	for _, c := range cases {
		if join(c.s, c.words...) != c.expected {
			t.Errorf("error: s: %s\tword: %v\texpected: %s\n", c.s, c.words, c.expected)
		}
	}
}
