package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//conn, err := net.Dial("tcp", "localhost:8888")
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal("tcp conn error")
	}
	defer conn.Close()
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	tcpConn.CloseWrite()
	<-done
}
