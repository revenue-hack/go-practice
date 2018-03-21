package ex02

import (
	"bytes"
	"fmt"
)

const UINT_SIZE = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/UINT_SIZE, uint(x%UINT_SIZE)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	//fmt.Printf("start: %v\n",x)
	word, bit := x/UINT_SIZE, uint(x%UINT_SIZE)
	//fmt.Printf("%v\t%v\t%v\t%v\n", word, bit, s.words, len(s.words))
	for word >= len(s.words) {
		//fmt.Printf("%v\t%v\n", word, s.words)
		s.words = append(s.words, 0)
	}
	// 1<<bit 2のbit乗をs.words[word]に足す.よって0+2^bit
	s.words[word] |= 1 << bit
	//fmt.Printf("%v\t%v\t%v\t%v\n", word, s.words, s.words[word], 1<<bit)
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UINT_SIZE; j++ {
			// 1<<uint(j)=2の累乗で増えていく(2,4,8,16...)
			// word&(1<<uint(j))=wordの1が立っている数を探す.ex)10&10=2
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", UINT_SIZE&i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.words = []uint{}
}

func (s *IntSet) Copy() *IntSet {
	var s2 IntSet
	s2.words = make([]uint, len(s.words))
	copy(s2.words, s.words)
	return &s2
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] &^= word
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] ^= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] &= word
		}
	}
	for i := len(t.words); i < len(s.words); i++ {
		s.words[i] = 0
	}
}

func (s *IntSet) Len() int {
	length := 0
	for _, word := range s.words {
		length += bitCount(word)
	}
	return length
}

func bitCount64(i uint64) int {
	i = i - ((i >> 1) & 0x5555555555555555)
	i = (i & 0x3333333333333333) + ((i >> 2) & 0x3333333333333333)
	i = (i + (i >> 4)) & 0x0f0f0f0f0f0f0f0f
	i = i + (i >> 8)
	i = i + (i >> 16)
	i = i + (i >> 32)
	return int(i) & 0x6f
}

func bitCount32(i uint32) int {
	i = i - ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	i = (i + (i >> 4)) & 0x0f0f0f0f
	i = i + (i >> 8)
	i = i + (i >> 16)
	i = i + (i >> 32)
	return int(i) & 0x6f
}

func bitCount(i uint) int {
	if UINT_SIZE == 32 {
		return bitCount32(uint32(i))
	} else {
		return bitCount64(uint64(i))
	}
}
