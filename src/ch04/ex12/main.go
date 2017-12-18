package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Comics struct {
	Month      string
	Num        int
	Link       string
	Year       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Title      string
	Day        string
}

const url = "https://xkcd.com/%s/info.0.json"

func main() {

	downloadComics(571)
}

func downloadComics(num int) {
	url := fmt.Sprintf(url, strconv.Itoa(num))
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("downloadComics response Error")
		panic(err)
	}
	if response.StatusCode != http.StatusOK {
		fmt.Printf("response statusCode Error %v in downloadComics\n", response.StatusCode)
		os.Exit(1)
	}
	var comics Comics
	if err := json.NewDecoder(response.Body).Decode(&comics); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", comics)
	fmt.Printf("%v\n", comics.Title)

}
