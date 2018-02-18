package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

var sign chan struct{}

func handleConn(c net.Conn) {
	sign = make(chan struct{}, 1)
	input := bufio.NewScanner(c)
	go selectWatch(c)
	for input.Scan() {
		sign <- struct{}{}
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func selectWatch(c net.Conn) {
	for {
		select {
		case <-time.After(10 * time.Second):
			c.Close()
		case <-sign:
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
