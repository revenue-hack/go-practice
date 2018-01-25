package main

import (
	"testing"
)

func TestIntSet_IntersectWith(t *testing.T) {
	for _, c := range []struct {
		c1       []int
		c2       []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2}, []int{1, 2}},
		{[]int{}, []int{1, 2}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5}, []int{1, 5}},
	} {
		var x IntSet
		var y IntSet
		x.AddAll(c.c1...)
		y.AddAll(c.c2...)
		x.IntersectWith(&y)
		for _, e := range c.expected {
			if !x.Has(e) {
				t.Errorf("intersectWith don't match %d\t%v\n", e, x)
			}
		}
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	for _, c := range []struct {
		c1       []int
		c2       []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2}, []int{3, 4, 5}},
		{[]int{}, []int{1, 2}, []int{}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5}, []int{2, 3, 4}},
	} {
		var x IntSet
		var y IntSet
		x.AddAll(c.c1...)
		y.AddAll(c.c2...)
		x.DifferenceWith(&y)
		if x.Len() != len(c.expected) {
			t.Errorf("differenceWith don't match %v\n", x)
		}
	}
}

func TestIntSet_SymmetricDifference(t *testing.T) {
	for _, c := range []struct {
		c1       []int
		c2       []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2}, []int{3, 4, 5}},
		{[]int{}, []int{1, 2}, []int{1, 2}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},

		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}, []int{1, 2, 3, 6, 7, 8}},
	} {
		var x IntSet
		var y IntSet
		x.AddAll(c.c1...)
		y.AddAll(c.c2...)
		x.SymmetricDifference(&y)
		if x.Len() != len(c.expected) {
			t.Errorf("symmetricDifference Len don't match %v\n", x)
		}
		for i, e := range c.expected {
			if !x.Has(e) {
				t.Errorf("symmetricDifference Has don't match %d\t%v\n", i, x)
			}
		}
	}
}

func (s *IntSet) Len() int {
	length := 0
	for _, word := range s.words {
		length += bitCount(word)
	}
	return length
}

func bitCount(i uint64) int {
	i = i - ((i >> 1) & 0x5555555555555555)
	i = (i & 0x3333333333333333) + ((i >> 2) & 0x3333333333333333)
	i = (i + (i >> 4)) & 0x0f0f0f0f0f0f0f0f
	i = i + (i >> 8)
	i = i + (i >> 16)
	i = i + (i >> 32)
	return int(i) & 0x6f
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
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
