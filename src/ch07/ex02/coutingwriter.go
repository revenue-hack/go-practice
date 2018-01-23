package main

import (
	"io"
	"fmt"
	"os"
)

func main() {
	bytes := []byte("あいうえお かかかかか hohoho　あぁぁｌ\nhoge hgoe\naaaa")
	w, c := CountingWriter(os.Stdout)
	n, err := w.Write(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%d\t%d\n", n, c)
}

type CWriter struct {
	cw int64
	cwIo io.Writer
}


func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cw CWriter
	cw.cwIo = w
	return &cw, &cw.cw
}

func (cw *CWriter) Write(p []byte) (int, error) {
	n, err := cw.cwIo.Write(p)
	cw.cw += int64(n)
	return n, err
}
