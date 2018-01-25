package main

import "fmt"

func main() {
	var x IntSet
	x.Add(1)
	x.Add(222)
	x.Add(3)
	fmt.Printf("value: %d\thas: %v\n", 222, x.Has(222))
	x.AddAll(1, 12)
	fmt.Printf("value: %d\thas: %v\n", 12, x.Has(12))
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

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}
