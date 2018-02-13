package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	go f("hello")
	go func() {
		log.Println("go func")
		log.Println(runtime.NumGoroutine())
	}()
	time.Sleep(time.Second)
}

func f(msg string) {
	log.Println(msg)
}
