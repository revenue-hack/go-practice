package github

import (
	"time"
)

const issuesURL = "https://api.github.com/repos/revenue-hack/go-practice/issues"
const contentType = "application/vnd.github.v3.full+json"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
}

type Issue struct {
	Number    int
	Title     string
	State     string
	CreatedAt time.Time `json:"created_at"`
}

type Credential struct {
	user string
	password string
}

type CreateIssue struct {
	Title string `json:"title"`
	Body string `json:"body"`
}

type CloseIssue struct {
	State string `json:"state"`
}

