package main

import "testing"

var dummyprereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"computer organization": {"compilers"},

	"data structures":  {"discrete math"},
	"databases":        {"data structures"},
	"discrete math":    {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks":         {"operating systems"},
	"operating systems": {
		"data structures",
		"computer organization",
	},
	"programming languages": {
		"data structures",
		"computer organization",
	},
}

var cases = []struct{
	maps map[string][]string
	expected bool
}{
	{maps: prereqs, expected: true},
	{maps: dummyprereqs, expected: false},
}

func TestTopoSort(t *testing.T) {
	for _, c := range cases {
		is := true
		maps := make(map[string]int)
		result := topoSort(c.maps)
		for j, r := range result {
			maps[r] = j
		}
		for m, i := range maps {
			for _, prereq := range c.maps[m] {
				if i < maps[prereq] {
					is = false
				}
			}
		}
		if c.expected != is {
			t.Errorf("don't match")
		}
	}
}
