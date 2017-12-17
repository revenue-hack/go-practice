package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	var shaFlag int
	flag.IntVar(&shaFlag, "sha", 256, "sha flag")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("command example: go run main.go -sha=384 value")
		return
	}
	var value string
	if len(os.Args) == 2 {
		value = os.Args[1]
	} else if len(os.Args) == 3 {
		value = os.Args[2]
	}
	switch shaFlag {
	case 384:
		fmt.Printf("384: %x\n", sha512.Sum384([]byte(value)))
	case 512:
		fmt.Printf("512: %x\n", sha512.Sum512([]byte(value)))
	default:
		fmt.Printf("256: %x\n", sha256.Sum256([]byte(value)))
	}
}
