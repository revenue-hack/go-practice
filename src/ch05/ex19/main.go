package main

import "fmt"

func main() {
	fmt.Printf("%d\n", variable(10))
}

func variable(v int) (variable int) {
	defer func() {
		if r := recover(); r != nil {
			variable = r.(int)
		}
	}()
	panic(v)
}
