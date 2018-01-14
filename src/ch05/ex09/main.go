package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("result: %s\n", expand(os.Args[1], func(s string) string {
		return strings.ToUpper(s)
	}))
}

func expand(s string, f func(string) string) string {
	reg, err := regexp.Compile(`\$\w+`)
	if err != nil {
		panic(err)
	}
	return reg.ReplaceAllStringFunc(s, func(s string) string {
		return f(s[1:])
	})
}
