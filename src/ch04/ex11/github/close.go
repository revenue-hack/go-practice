package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Close(issueId int) bool {
	jsonBody := CloseIssue{"closed"}
	body, err := json.Marshal(&jsonBody)
	if err != nil {
		panic(err)
	}
	url := issuesURL + "/" + strconv.Itoa(issueId)
	response, err := patch(url, bytes.NewReader(body))
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("close response error: %v\n", err)
		return false
	}
	if response.StatusCode != http.StatusOK {
		fmt.Printf("close response status not 200: %v\n", response)
		return false
	}
	return true
}
