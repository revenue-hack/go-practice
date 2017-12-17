package github

import (
	"fmt"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

func Read(issueId int) Issue {
	url := issuesURL + "/" + strconv.Itoa(issueId)
	response, err := get(url)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("%v\n", err)
		panic(err)
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("read response status not 200")
		os.Exit(1)
	}
	var issues Issue
	if err := json.NewDecoder(response.Body).Decode(&issues); err != nil {
		fmt.Println("read response decode error")
		panic(err)
	}
	return issues
}

