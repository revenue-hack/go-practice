package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	clocks, err := parseArgs()
	if err != nil {
		panic(err)
	}
	ch := make(chan string, 3)
	for _, clock := range clocks {
		go func(clock *Clock) {
			conn, err := net.Dial("tcp", clock.connection)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			reader := bufio.NewReader(conn)
			for {
				bytes, _, err := reader.ReadLine()
				if err != nil {
					return
				}
				ch <- fmt.Sprintf("%s: %v", clock.location, string(bytes))
			}
		}(clock)
	}
	for {
		if num := len(ch); num == 3 {
			var str []string
			for i := 0; i < num; i++ {
				str = append(str, <-ch)
			}
			fmt.Println(strings.Join(str, " "))
		}
	}
}

type Clock struct {
	location   string
	connection string
}

func parseArgs() ([]*Clock, error) {
	if len(os.Args) == 0 {
		return nil, fmt.Errorf("fatal args")
	}
	var clocks []*Clock
	for _, arg := range os.Args[1:] {
		strs := strings.Split(arg, "=")
		if len(strs) != 2 {
			return nil, fmt.Errorf("illegal args")
		}
		clocks = append(clocks, &Clock{location: strs[0], connection: strs[1]})
	}
	return clocks, nil
}
