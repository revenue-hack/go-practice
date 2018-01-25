package main

import (
	"sort"
	"fmt"
)

func main() {
	s1 := "hhhh"
	b1 := []byte(s1)
	s2 := "hello"
	b2 := []byte(s2)
	fmt.Printf("%s: %v\t%s: %v\n", s1, IsPalindrome(Byte(b1)), s2, IsPalindrome(Byte(b2)))
}

func IsPalindrome(s sort.Interface) bool {
	length := s.Len()
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		} else {
			return false
		}
	}
	return true
}

type Byte []byte

func (b Byte) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b Byte) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b Byte) Len() int {
	return len(b)
}
