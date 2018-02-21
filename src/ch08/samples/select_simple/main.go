package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int, 1)
	/*
		for i := 0; i < 10; i++ {
			select {
			case x := <-ch:
				log.Println(x)
			case ch <- i:
			}
		}
	*/
	for i := 0; ; i++ {
		select {
		case <-time.After(2 * time.Second):
			log.Println("sum")
			return
		case x := <-ch:
			log.Println(x)
		case ch <- i:
		}
	}

}
