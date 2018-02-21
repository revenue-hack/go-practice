package main

import (
	"log"
	"sync"
	"time"
)

// これは動く
func main() {
	done := make(chan struct{}, 0)
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	timeoutCh := make(chan struct{})
	var mu sync.Mutex
	var counter int
	go func() {
		for {
			select {
			case <-timeoutCh:
				close(done)
				return
			case <-ch1:
				mu.Lock()
				counter++
				mu.Unlock()
				ch2 <- struct{}{}
			}
		}
	}()
	go func() {
		for {
			select {
			case <-timeoutCh:
				close(done)
				return
			case <-ch2:
				mu.Lock()
				counter++
				mu.Unlock()
				ch1 <- struct{}{}
			}
		}
	}()
	ch1 <- struct{}{}
	select {
	case <-time.After(1 * time.Second):
		timeoutCh <- struct{}{}
	}
	<-done
	log.Printf("result: %d\n", counter)

}
