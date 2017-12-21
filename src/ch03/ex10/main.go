package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", comma("djaljvoajkdfjdoadwjfdoidjw"))
}

func comma(s string) string {
	sBytes := []byte(s)
	if len(sBytes) < 3 {
		return s
	}
	var buf bytes.Buffer
	remainder := len(sBytes) % 3
	if remainder == 0 {
		remainder = 3
	}
	for _, v := range sBytes {
		if remainder == 0 {
			buf.WriteString(",")
			remainder = 3
		}
		buf.WriteByte(v)
		remainder--
	}
	return buf.String()
}
