package main

import (
	"fmt"
	"sync"
	"time"
)

var start = time.Now()
var counter int
var mu sync.Mutex

const GOROUTINE_MAX_COUNT = 10000000

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	var total int64
	go func() {
		for x := 0; ; x++ {
			start = time.Now()
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		mu.Lock()
		counter++
		mu.Unlock()
		<-squares
		t := time.Since(start)
		fmt.Printf("goroutine num: %d\ttime: %v\n", counter, t)
		total += t.Nanoseconds()
		if counter > GOROUTINE_MAX_COUNT {
			break
		}
	}
	fmt.Printf("average: %v ns\n", total/GOROUTINE_MAX_COUNT)

}

// goroutine num: 9999991	time: 270ns
// goroutine num: 9999992	time: 2.203µs
// goroutine num: 9999993	time: 4.723µs
// goroutine num: 9999994	time: 239ns
// goroutine num: 9999995	time: 2.382µs
// goroutine num: 9999996	time: 5.046µs
// goroutine num: 9999997	time: 230ns
// goroutine num: 9999998	time: 2.557µs
// goroutine num: 9999999	time: 5.499µs
// goroutine num: 10000000	time: 239ns
// goroutine num: 10000001	time: 2.454µs
// average: 4041 ns
