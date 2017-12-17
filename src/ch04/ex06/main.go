package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	//strs := "asb"
	//fmt.Printf("%v\n", zip(strs))
}

func zip(strs string) {
	spaceBuf := make([]byte, 4)
	spaceSize := utf8.EncodeRune(spaceBuf, ' ')
	spaceBuf = spaceBuf[:spaceSize]
	/*
	strBytes := []byte(strs)
	for i, byte := range strBytes {
		strRune, _ := utf8.DecodeRune(byte)
		if unicode.IsSpace(byte) {

		}

	}
	for i, str := range strs {
		// str is rune
		if unicode.IsSpace(str) {

			copy(, spaceBuf)

		}
	}
	return strs
	*/
}
