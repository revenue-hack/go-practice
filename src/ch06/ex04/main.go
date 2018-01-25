package main

import (
	"bytes"
	"fmt"
)

func main() {
	var x IntSet
	x.Add(1)
	x.Add(1111)
	x.Add(234)
	fmt.Printf("%v\n", x.Elems())
}

type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	//fmt.Printf("start: %v\n",x)
	word, bit := x/64, uint(x%64)
	//fmt.Printf("%v\t%v\t%v\t%v\n", word, bit, s.words, len(s.words))
	for word >= len(s.words) {
		//fmt.Printf("%v\t%v\n", word, s.words)
		s.words = append(s.words, 0)
	}
	// 1<<bit 2のbit乗をs.words[word]に足す.よって0+2^bit
	s.words[word] |= 1 << bit
	//fmt.Printf("%v\t%v\t%v\t%v\n", word, s.words, s.words[word], 1<<bit)
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			// 1<<uint(j)=2の累乗で増えていく(2,4,8,16...)
			// word&(1<<uint(j))=wordの1が立っている数を探す.ex)10&10=2
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64&i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
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

	elems := make([]int, 0, len)
	for i, sword := range s.words {
		for bit := uint(0); bit < 64; bit++ {
			if sword&(1<<bit) != 0 {
				elems = append(elems, i*64+int(bit))
			}
		}
	}
	return elems
}
