package main

import "log"

func main() {
	ch := make(chan bool)
	go f(ch)
	log.Println(<-ch)

}

func f(ch chan bool) {
	ch <- true
}
