package main

import (
	"log"
	"time"
)

// これは動く
func main() {
	done := make(chan struct{})
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ch1:
				log.Print("ch1")
				ch2 <- struct{}{}
			}
		}
	}()
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ch2:
				log.Print("ch2")
				ch1 <- struct{}{}
			}
		}
	}()
	ch1 <- struct{}{}
	time.Sleep(2 * time.Second)
	done <- struct{}{}
	log.Printf("result: %d\n", 10)

}
