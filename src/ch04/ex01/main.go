package main

import (
	"crypto/sha256"
	"fmt"
)

var bytes [256]byte

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%v\n", compareSha256(c1, c2))
}

func init() {
	for i := range bytes {
		bytes[i] = bytes[i/2] + byte(i&1)
	}
}

func compareSha256(c1, c2 [32]byte) int {
	var sum int
	for i, v1 := range c1 {
		sum += int(bytes[v1^c2[i]])
	}
	return sum
}
