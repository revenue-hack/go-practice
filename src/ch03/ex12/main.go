package main

import (
	"fmt"
	"sort"
	"bytes"
)

type Byte []byte

func main() {
	fmt.Printf("%v\n", anagram("hlleo", "hello"))
}

func anagram(s1, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
	sort.Sort(Byte(b1))
	sort.Sort(Byte(b2))
	ans := bytes.Compare(b1, b2)
	if ans == 0 { return true } else { return false }
}

func (b Byte) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b Byte) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b Byte) Len() int {
	return len(b)
}