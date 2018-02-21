package main

import (
	"log"
	"time"
)

func main() {
	done := make(chan struct{}, 0)
	end := make(chan struct{})
	countCh := make(chan int)
	go counter(countCh, end)
	go func() {
		select {
		case <-time.After(5 * time.Second):
			log.Println("after")
			end <- struct{}{}
			close(done)
		}
	}()

	<-done
	log.Println("done")
	log.Printf("result: %d\n", <-countCh)
}

// deadlockする
/*
func counter(countCh chan int, end <-chan struct{}) {
	for i := 0; ; i++ {
		select {
		case <-end:
			break
		default:
			countCh <- i
			log.Print(i)
		}
	}
}
*/

// これならちゃんとカウントされる
func counter(countCh chan int, end <-chan struct{}) {
	for i := 0; ; i++ {
		select {
		case <-end:
			countCh <- i
			break
		default:
		}
	}
}
