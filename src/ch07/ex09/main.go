package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
)

var title = func(x, y *Track) bool {
	return x.Title < y.Title
}
var year = func(x, y *Track) bool {
	return x.Year < y.Year
}
var artist = func(x, y *Track) bool {
	return x.Artist < y.Artist
}
var album = func(x, y *Track) bool {
	return x.Album < y.Album
}
var leng = func(x, y *Track) bool {
	return x.Length < y.Length
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := createSortData(r)
	w.Header().Set("Content-Type", "text/html")
	printTraceHtml().Execute(w, response(data))
}

func createSortData(r *http.Request) []*Track {
	data := createData()
	sortTable := MainSort{t: data}
	if err := r.ParseForm(); err != nil {
		fmt.Printf("%v\n", "hohhhhhhhhhhhhhhh")
		return data
	}
	if len(r.Form) != 0 {
		for k, _ := range r.Form {
			switch k {
			case "title":
				sortTable.addSortKey(title)
			case "artist":
				sortTable.addSortKey(artist)
			case "album":
				sortTable.addSortKey(album)
			case "year":
				sortTable.addSortKey(year)
			case "length":
				sortTable.addSortKey(leng)
			default:
				sortTable.addSortKey(title)
			}
		}
		sort.Sort(&sortTable)
	}
	return data
}

func response(tracks []*Track) Response {
	return Response{tracks}
}

type Response struct {
	TrackData []*Track
}

func createData() []*Track {
	data := make([]*Track, len(tracksData))
	copy(data, tracksData)
	return data
}

func printTraceHtml() *template.Template {
	return template.Must(template.New("response").Parse(`
		<h2>sing list</h2>
		<table border='1' width='80%'>
			<tr>
				<th>Title</th>
				<th>Artist</th>
				<th>Album</th>
				<th>Year</th>
				<th>Length</th>
			</tr>
			{{range .TrackData}}
			<tr>
				<td>{{.Title}}</td>
				<td>{{.Artist}}</td>
				<td>{{.Album}}</td>
				<td>{{.Year}}</td>
				<td>{{.Length}}</td>
			</tr>
			{{end}}
		</table>
	`))
}
