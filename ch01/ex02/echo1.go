package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("%s", echo(i, arg))
	}
}

func echo(i int, s string) string {
	return fmt.Sprintf("i: %d value: %s\n", i, s)
}
