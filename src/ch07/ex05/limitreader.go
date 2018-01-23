package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	lr := LimitReader(strings.NewReader("abcdefghij"), int64(3))
	n, err := lr.Read(make([]byte, 10))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", n)
	iolr := io.LimitReader(strings.NewReader("abcdefghij"), int64(3))
	ioN, err := iolr.Read(make([]byte, 10))
	fmt.Printf("%v\n", ioN)
}

type limitedReader struct {
	r    io.Reader
	n    int64
	next int
}

func (lr *limitedReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	fmt.Printf("p: %v\n", p)
	if int64(lr.next) >= lr.n {
		return 0, io.EOF
	}
	bytes := int(lr.n - int64(lr.next))
	fmt.Printf("bytes: %v\n", bytes)
	if bytes > len(p) {
		bytes = len(p)
	}
	n, err := lr.r.Read(p[:bytes])
	lr.next += bytes
	fmt.Printf("next: %v\n", lr.next)
	return n, err

}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitedReader{r, n, 0}
}
