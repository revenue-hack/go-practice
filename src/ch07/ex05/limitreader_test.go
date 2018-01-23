package main

import (
	"io"
	"strings"
	"testing"
)

func TestLimitedReader_Read(t *testing.T) {
	var readers = []struct {
		str  string
		n    int64
		next int
	}{
		{"aaaaaaaaaaaa", 3, 0},
		{"aaaaaaaaaaaa", 30, 0},
		{"", 30, 0},
	}
	for _, reader := range readers {
		lr := LimitReader(strings.NewReader(reader.str), reader.n)
		iolr := io.LimitReader(strings.NewReader(reader.str), reader.n)
		for i := 0; i < 2; i++ {
			n, err := lr.Read(make([]byte, 30))
			ioN, errN := iolr.Read(make([]byte, 30))
			if n != ioN {
				t.Errorf("n don't match io: %v\tlimit: %v\n", ioN, n)
			}
			if err != errN {
				t.Errorf("err don't match io: %v\tlimit: %v\n", errN, err)
			}
		}
	}
}
