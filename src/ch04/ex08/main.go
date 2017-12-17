package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type Types struct {
	Digit  int
	Letter int
	Lower  int
	Space  int
	Title  int
	Upper  int
}

var summary Types

func countType(rune int32) {
	switch {
	case unicode.IsDigit(int32(rune)):
		summary.Digit++
	case unicode.IsLetter(int32(rune)):
		summary.Letter++
	case unicode.IsLower(int32(rune)):
		summary.Lower++
	case unicode.IsSpace(int32(rune)):
		summary.Space++
	case unicode.IsTitle(int32(rune)):
		summary.Title++
	case unicode.IsUpper(int32(rune)):
		summary.Upper++
	}
}

func dispType() {
	fmt.Printf("\ntype\tcount\n")
	fmt.Printf("Digit\t%d\n", summary.Digit)
	fmt.Printf("Letter\t%d\n", summary.Letter)
	fmt.Printf("Lower\t%d\n", summary.Lower)
	fmt.Printf("Space\t%d\n", summary.Space)
	fmt.Printf("Title\t%d\n", summary.Title)
	fmt.Printf("Upper\t%d\n", summary.Upper)
	fmt.Println()
}

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		countType(r)
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	dispType()
}
