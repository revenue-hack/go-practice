package main

import (
	"net/http"
	"log"
	"github.com/revenue-hack/go-practice/src/ch04/ex14/github"
	"html/template"
)

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	requestTemplate().Execute(w, response())
}

type Templates struct {
	Users []*github.User
	Bugs []*github.Bug
	Milestones []*github.Milestone
}

func response() Templates {
	return Templates{
		github.UsersRequest().Users,
		github.BugReportRequest().Bugs,
		github.MilestonesRequest().Milestones}
}

func requestTemplate() *template.Template {
	return template.Must(template.New("template").Parse(`
		<h2>User List</h2>
		<table border='1' width='80%'>
			<tr>
				<th>Login</th>
				<th>Id</th>
				<th>Url</th>
			</tr>
			{{range .Users}}
			<tr>
				<td>{{.Login}}</td>
				<td>{{.Id}}</td>
				<td>{{.Url}}</td>
			</tr>
			{{end}}
		</table>
		<h2>Milestone</h2>
		<table border='1' width='80%'>
			<tr>
				<th>Number</th>
				<th>Title</th>
				<th>Description</th>
			</tr>
			{{range .Milestones}}
			<tr>
				<td>{{.Number}}</td>
				<td>{{.Title}}</td>
				<td>{{.Description}}</td>
			</tr>
			{{end}}
		</table>
		<h2>Bug Report</h2>
		<table border='1' width='80%'>
			<tr>
				<th>Number</th>
				<th>Title</th>
				<th>Body</th>
				<th>Url</th>
			</tr>
			{{range .Bugs}}
			<tr>
				<td>{{.Number}}</td>
				<td>{{.Title}}</td>
				<td>{{.Body}}</td>
				<td>{{.Url}}</td>
			</tr>
			{{end}}
		</table>
	`))

}
