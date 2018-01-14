package main

import "fmt"

func main() {
	var x IntSet
	x.words = []uint64{1,1111, 234}
	fmt.Printf("%v\n", x.Elems())
}

type IntSet struct {
	words []uint64
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
	return int(i) & 0x7f
}

func (s *IntSet) Elems() []int {
	len := s.Len()
	if len == 0 {
		return []int{}
	}

	enums := make([]int, 0, len)
	for i, sword := range s.words {
		for bit := uint(0); bit < 64; bit++ {
			if sword&(1<<bit) != 0 {
				enums = append(enums, i*64+int(bit))
			}
		}
	}
	return enums
}