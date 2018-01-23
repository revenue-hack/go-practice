package main

import (
	"testing"
	"os"
)

var cases = []string{
	"hoge\n",
	"あいうえお\n",
	"dwhoge aaaa\n",
	"あああ　ああほ\n",
}

func TestCountingWriter(t *testing.T) {
	w, count := CountingWriter(os.Stdout)
	var total int64
	for _, c := range cases {
		bytes := []byte(c)
		w.Write(bytes)
		total += int64(len(bytes))
		if *count != total {
			t.Errorf("[]couting miss] count %d\ttotal %d\n", count, total)
		}
	}
}
