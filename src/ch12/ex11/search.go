package main

import (
	"ch12/ex11/params"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}

type NoTagData struct {
	Labels     []string
	MaxResults int
	Exact      bool
}

func search(resp http.ResponseWriter, req *http.Request) {
	var data Data
	data.MaxResults = 10
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(resp, "Search: %+v\n", data)
}

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
