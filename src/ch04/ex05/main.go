package main

import (
	"fmt"
)

func main() {
	strs := [...]string{"a", "a", "ss", "2"}
	fmt.Printf("%v\n", unique(strs[:]))
}

func unique(strs []string) []string {
	var i int
	for i != len(strs)-1 {
		if strs[i] == strs[i+1] {
			strs = remove(strs, i)
			i = 0
			continue
		}
		i++
	}
	return strs
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
