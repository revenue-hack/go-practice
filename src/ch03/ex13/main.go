package main

import "fmt"

const (
	KiB = 1024
	MiB = 1024 * 1024
	GiB = MiB * 1024
	TiB = GiB * 1024
	PiB = TiB * 1024
	ZiB = PiB * 1024
	YiB = ZiB * 1024
)

func main() {
	fmt.Printf("%v\n", KiB)
	fmt.Printf("%v\n", MiB)
	fmt.Printf("%v\n", GiB)
	fmt.Printf("%v\n", TiB)
	fmt.Printf("%v\n", PiB)
	fmt.Printf("%v\n", ZiB)
	// overflow
	fmt.Printf("%v\n", YiB)
}
