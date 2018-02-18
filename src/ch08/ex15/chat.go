package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	who string
	ch  chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				select {
				case cli.ch <- msg:
				default:
				}
			}
		case cli := <-entering:
			for client := range clients {
				cli.ch <- client.who + " is arrive"
			}
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}

}

const TIME = 5 * time.Minute

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	input := bufio.NewScanner(conn)
	// first input
	if input.Scan() {
		text := input.Text()
		r := regexp.MustCompile(`<user>(.+)</user>`)
		if r.Match([]byte(text)) {
			result := r.FindAllStringSubmatch(text, -1)
			who = result[0][1]
		}
	}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{who, ch}
	timer := time.AfterFunc(TIME, func() {
		conn.Close()
	})

	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer = time.AfterFunc(TIME, func() {
			conn.Close()
		})
	}
	timer.Stop()
	leaving <- client{who, ch}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
