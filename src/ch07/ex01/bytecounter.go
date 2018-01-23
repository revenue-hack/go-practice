package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	bytes := []byte("あいうえお かかかかか hohoho　あぁぁｌ\nhoge hgoe\naaaa")
	var w WordCounter
	_, err := w.Write(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wordCounter: %d\n", w)

	var l LineCounter
	_, err = l.Write(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("lineCounter: %d\n", l)
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	for scanner.Scan() {
		_ = scanner.Text()
		*l += 1
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	b := p
	var advanceCount int
	for {
		advance, token, err := bufio.ScanWords(b, true)
		if err != nil {
			panic(err)
		}
		if token != nil {
			*w += 1
		}
		advanceCount += advance
		b = b[advance:]
		if advanceCount == len(p) {
			return len(p), nil
		}
	}
}
