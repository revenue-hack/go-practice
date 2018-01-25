package main

import "testing"

func TestElems(t *testing.T) {
	for _, c := range []struct {
		c1 []int
	}{
		{[]int{}},
		{[]int{1, 10, 100, 1000, 10000}},
		{[]int{1}},
		{[]int{1, 2, 3, 4, 5}},
	} {
		var x IntSet
		x.AddAll(c.c1...)

		elems := x.Elems()
		if len(elems) != len(c.c1) {
			t.Errorf("Elems: %d\n", len(c.c1))
		}

		for i, value := range elems {
			if !x.Has(value) {
				t.Errorf("Elems: i: %d\t%value: d\n", i, value)
			}
		}
	}
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

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}
