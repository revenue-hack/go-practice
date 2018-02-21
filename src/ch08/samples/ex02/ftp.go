package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

var (
	baseDir = os.Getenv("HOME")
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
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

func handleConn(conn net.Conn) {
	fmt.Println("FTP Server Connect")
	input := bufio.NewScanner(conn)
	for input.Scan() {
		if input.Text() == "" {
			continue
		}
		cmds := strings.Split(input.Text(), " ")
		switch cmds[0] {
		case "ls":
			List()
		case "cd":
			var target string
			if len(cmds) != 1 {
				target = cmds[1]
			}
			dir, err := ChangeDir(target)
			if err != nil {
				fmt.Printf("cd: no such file or directory: %s\n", dir)
			} else {
				baseDir = dir
			}
			fmt.Printf("current pwd: %s\n", baseDir)
		default:
			fmt.Printf("invalid command: %s\n", cmds[0])
		}
	}
}

func List() {
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fileType := "f"
		if file.IsDir() {
			fileType = "d"
		}
		fmt.Printf("%s\t%d\t%s\n", fileType, file.Size(), file.Name())
	}
}

// @todo ../とか対応してない
func ChangeDir(target string) (string, error) {
	var dir string
	if target == "" {
		dir = os.Getenv("HOME")
	} else if strings.HasPrefix(target, "/") {
		dir = target
	} else if strings.HasPrefix(target, "./") {
		dir = fmt.Sprintf("%s/%s", baseDir, string([]rune(target[2:])))
	} else {
		dir = fmt.Sprintf("%s/%s", baseDir, target)
	}
	return dir, os.Chdir(dir)
}

func Close(conn net.Conn) error {
	conn.Close()

}
