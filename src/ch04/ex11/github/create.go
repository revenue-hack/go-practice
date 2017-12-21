package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Create(title, contents string) bool {
	jsonBody := CreateIssue{title, contents}
	body, err := json.Marshal(&jsonBody)
	fmt.Printf("%v\n", jsonBody)
	if err != nil {
		panic(err)
	}
	response, err := post(issuesURL, bytes.NewReader(body))
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("create response error: %v\n", err)
		return false
	}
	if response.StatusCode != http.StatusCreated {
		fmt.Printf("create response status not 201: %v\n", response)
		return false
	}
	return true
}
