package github

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
)

const bugUrl = "https://api.github.com/repos/revenue-hack/go-practice/issues"

type Bug struct {
	Url string
	Number int
	Title string
	Body string
}

type BugList struct {
	Bugs []*Bug
}

func BugReportRequest() BugList {
	response, err := get(bugUrl)
	if err != nil {
		fmt.Println("BugReportRequest response panic")
		panic(err)
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("read response status not 200")
		os.Exit(1)
	}
	var bugList BugList
	if err := json.NewDecoder(response.Body).Decode(&(bugList.Bugs)); err != nil {
		fmt.Println("read response decode error")
		panic(err)
	}
	return bugList
}
