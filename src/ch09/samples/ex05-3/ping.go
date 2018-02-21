package main

import (
	"log"
	"time"
)

func main() {
	done := make(chan struct{})
	countCh := make(chan int)
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go counter(countCh, done, ch1, ch2)
	go counter(countCh, done, ch1, ch2)
	go func(ch1 chan struct{}) {
		select {
		case <-time.After(5 * time.Second):
			log.Println("after")
			done <- struct{}{}
		default:
			ch1 <- struct{}{}
		}
	}(ch1)
	log.Printf("result: %d\n", <-countCh)
}

// これならちゃんとカウントされる
/*
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
*/
// これは出来ない
func counter(countCh chan int, end <-chan struct{}, ch1 chan struct{}, ch2 chan struct{}) {
	for i := 0; ; i++ {
		select {
		case <-end:
			log.Println("end")
			countCh <- i
			break
		case <-ch1:
			ch2 <- struct{}{}
		case <-ch2:
			ch1 <- struct{}{}
		}
	}
}
