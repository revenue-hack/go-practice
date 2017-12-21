package main

import (
	"fmt"
)

func main() {
	bytes := []byte("asbyuio")
	fmt.Printf("%v\n", bytes)
	fmt.Printf("%v\n", reverse(bytes))
}

func reverse(bytes []byte) []byte {
	variableMax := len(bytes) - 1
	halfLength := len(bytes) / 2
	for i, byte := range bytes {
		bytes[i] = bytes[variableMax]
		bytes[variableMax] = byte
		variableMax--
		if variableMax < halfLength {
			break
		}
	}
	return bytes
}
