package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	var port = flag.Int("port", 8888, "port number")
	flag.Parse()

	address := "localhost:" + strconv.Itoa(*port)
	listener, err := net.Listen("tcp", address)
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

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("10:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
