package main

import (
	"net/http"
	"log"
	"fmt"
	"os"
	"encoding/json"
	"io"
	"html/template"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/users", userHandler)
	http.HandleFunc("/milestone", milestoneHandler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Printf("%v\n", r)
	userTemplate().Execute(w, usersRequest())
}

const url = "https://api.github.com/users"

type User struct {
	Login string
	Id int
	Url string
	HtmlUrl string `json:"html_url"`
}

type UserList struct {
	Users []*User
}

func usersRequest() UserList {
	response, err := get(url)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("usersRequest response panic")
		panic(err)
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("read response status not 200")
		os.Exit(1)
	}
	var userList UserList
	if err := json.NewDecoder(response.Body).Decode(&(userList.Users)); err != nil {
		fmt.Println("read response decode error")
		panic(err)
	}
	return userList
}

func userTemplate() *template.Template {
	return template.Must(template.New("userList").Parse(`
		<h1>hogeho</h1>
		{{range .Users}}
		<h1>{{.Id}}</h1>
		{{end}}
	`))
}


func createRequest(method, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	return request
}

func get(url string) (*http.Response, error) {
	request := createRequest("GET", url, nil)
	return http.DefaultClient.Do(request)
}


func milestoneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

}
