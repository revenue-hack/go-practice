package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

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

func main() {
	ans := topoSort(prereqs)
	for i, course := range ans {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	fmt.Printf("%v\n", isTopologicalOrdered(ans))
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order
}

func isTopologicalOrdered(ts []string) error {
	maps := make(map[string]int)
	for i, course:= range ts {
		maps[course] = i
	}
	for course, i := range maps {
		for _, prereq := range prereqs[course] {
			if i < maps[prereq] {
				return fmt.Errorf("%s, %s cycle", course, prereq)
			}
		}
	}
	return nil
}
