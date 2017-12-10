package main

import (
	"fmt"
	"bytes"
)

func main() {
	fmt.Printf("%s\n", comma("djaljvoajkdfjdoadwjfdoidjw"))
}

func anagram(s1, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
	for _, v1 := range b1 {
		var match = false
		for i, v2 := range b2 {
			if v1 == v2 {
				b2 =
				b2 = bytes.Join(b2[0:i-1], b2[i])
				match = true
				break
			}
		}
		if
	}

}
