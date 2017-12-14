package main

import (
	"crypto/sha256"
	"fmt"
)

var bytes [256]byte

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	compareSha256(c1, c2)
}

func init() {
	for i := range bytes {
		bytes[i] = bytes[i/2] + byte(i&1)
	}
}

func compareSha256(c1, c2 [32]byte) int {
	fmt.Printf("%v\n", []byte("X"))
	fmt.Printf("%v\n", c1)
	fmt.Printf("%x\n", c1)
	fmt.Printf("%v\n", c2)
	fmt.Printf("hatto: %v\n", c1[0]^c2[0])
	var count int
	/*
	for i, v1 := range c1 {
		bytes[]
	}
	*/
	return count
}
