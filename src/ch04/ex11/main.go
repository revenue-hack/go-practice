package main

import (
	"flag"
	"fmt"
	"github.com/revenue-hack/go-practice/src/ch04/ex11/github"
)

func main() {
	var typeFlag, titleFlag, contentsFlag, stateFlag string
	var issueIdFlag int
	setFlag(&typeFlag, &titleFlag, &contentsFlag, &stateFlag, &issueIdFlag)

	switch typeFlag {
	case "create":
		create(titleFlag, contentsFlag)
	case "update":
		update(issueIdFlag, titleFlag, contentsFlag)
	case "read":
		read(issueIdFlag)
	case "close":
		close(issueIdFlag)
	}
}

func create(title, contents string) {
	if !github.Create(title, contents) {
		fmt.Println("create error")
	} else {
		fmt.Println("create ok")
	}
}

func update(id int, title, contents string) {
	if !github.Update(id, title, contents) {
		fmt.Println("update error")
	} else {
		fmt.Println("update ok")
	}
}

func read(id int) {
	issue := github.Read(id)
	fmt.Printf("%v\n", issue)
}

func close(id int) {
	if !github.Close(id) {
		fmt.Println("close error")
	} else {
		fmt.Println("close ok")
	}
}

func setFlag(typeFlag, title, contents, state *string, issueId *int) {
	flag.StringVar(typeFlag, "type", "none", "create or update or read or close")
	flag.StringVar(title, "title", "nontitle", "input issue title text")
	flag.StringVar(contents, "contents", "", "input issue body text")
	flag.StringVar(state, "state", "closed", "closed or open")
	flag.IntVar(issueId, "id", 1, "input issue id")
	flag.Parse()
}
