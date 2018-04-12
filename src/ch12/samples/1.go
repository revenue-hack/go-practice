package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}
	x = "100"
	switch x := x.(type) {
	case string:
		fmt.Printf("%v\n", x)
	default:
		fmt.Println("nothing")
	}

	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)
}
