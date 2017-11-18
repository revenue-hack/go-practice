package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	countFiles := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("nothing args")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts, countFiles, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, countFiles[line][0:])
		}
	}
	fmt.Printf("\n")
}

func countLines(f *os.File, counts map[string]int, countFiles map[string]string, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] != 0 && !strings.Contains(countFiles[input.Text()], fileName) {
			countFiles[input.Text()] += "," + fileName
		} else if counts[input.Text()] == 0 {
			countFiles[input.Text()] += fileName
		}
		counts[input.Text()]++
	}
}
