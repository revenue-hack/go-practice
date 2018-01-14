package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", join(", ", "sssssss", "aaaa", "hhhhhh"))

}

func join(s string, vals ...string) string {
	var result bytes.Buffer
	for i, val := range vals {
		result.WriteString(val)
		if i != len(vals)-1 {
			result.WriteString(s)
		}
	}
	return result.String()
}
