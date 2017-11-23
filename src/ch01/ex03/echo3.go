package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// inefficiency
	ine_s := time.Now()
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		fmt.Printf(inefficiencyEcho(arg))
	}
	fmt.Println("\n")
	ine_e := time.Since(ine_s)

	// efficiency
	e_s := time.Now()
	fmt.Printf(efficiencyEcho())
	fmt.Println("\n")
	e_e := time.Since(e_s)

	fmt.Printf("非効率: %.2fs 効率: %.2fs\n", ine_e, e_e)
}

func inefficiencyEcho(s string) string {
	return fmt.Sprintf("%s ", s)
}

func efficiencyEcho() string {
	return strings.Join(os.Args[1:], " ")
}
