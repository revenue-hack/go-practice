package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	result, sum := wordFreq(in)
	display(result, sum)
}

func display(result map[string]int, sum int) {
	fmt.Println()
	fmt.Printf("letter\tfrequency\n")
	for s, i := range result {
		fmt.Printf("%s\t%.2f\n", s, float64(i)/float64(sum)*100)
	}
}

func wordFreq(in *bufio.Scanner) (map[string]int, int) {
	result := make(map[string]int)
	in.Split(bufio.ScanWords)
	var sum int
	for in.Scan() {
		result[in.Text()]++
		sum++
	}
	return result, sum
}
