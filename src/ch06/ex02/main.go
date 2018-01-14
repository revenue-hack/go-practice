package main

import "fmt"

func main() {
	var x IntSet
	x.words = []uint64{1, 222, 2}
	fmt.Printf("%v\n", x)
	x.AddAll(1, 12)
	fmt.Printf("%v\n", x)
}


type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}
