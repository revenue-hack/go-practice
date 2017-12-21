package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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

const (
	url     = "https://xkcd.com/%s/info.0.json"
	maxPage = 1930
	fileDir = "/tmp/data"
)

func main() {
	if !isCache() {
		createCacheFile()
	}
	if len(os.Args) > 1 {
		s := os.Args[1]
		fmt.Printf("%sの検索結果\n", s)
		comicses := search(s)
		for _, comics := range comicses {
			fmt.Printf("%v\t%v\n", comics.Link, comics.Transcript)
		}
	}

}

func search(s string) []Comics {
	var comicses []Comics
	for i := 0; i < 10; i++ {
		page := i + 1
		if comics, isExist := readCacheFile(page); isExist {
			if strings.Contains(comics.Title, s) || strings.Contains(comics.Transcript, s) {
				comicses = append(comicses, comics)
			}
		}
	}
	return comicses
}

func isCache() bool {
	if _, dir := os.Stat(fileDir); dir != nil {
		return false
	} else {
		return true
	}
}

func createCacheFile() {
	for i := 0; i < maxPage; i++ {
		page := i + 1
		if _, err := os.Stat(fileDir); err != nil {
			os.Mkdir(fileDir, 0777)
		}
		filePath := fmt.Sprintf("%s/%d.json", fileDir, page)
		if _, err := os.Stat(filePath); err == nil {
			continue
		}
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("file create Error in createFileCache")
			panic(err)
		}
		defer file.Close()
		bytes, err := json.Marshal(downloadComics(page))
		if err != nil {
			fmt.Println("json parse Error in createFileCache")
			panic(err)
		}
		file.Write(([]byte)(bytes))
	}
}

func readCacheFile(page int) (Comics, bool) {
	filePath := fmt.Sprintf("%s/%d.json", fileDir, page)
	file, err := ioutil.ReadFile(filePath)
	var comics Comics
	if err != nil {
		return comics, false
	}
	if err := json.Unmarshal(file, &comics); err != nil {
		return comics, false
	}
	return comics, true
}

func downloadComics(page int) Comics {
	url := fmt.Sprintf(url, strconv.Itoa(page))
	response, err := http.Get(url)
	defer response.Body.Close()
	var comics Comics
	if err != nil {
		fmt.Printf("page:%d downloadComics response Error\n", page)
		return comics
	}
	if response.StatusCode != http.StatusOK {
		fmt.Printf("page:%d response statusCode Error %v in downloadComics\n", page, response.StatusCode)
		return comics
	}
	if err := json.NewDecoder(response.Body).Decode(&comics); err != nil {
		fmt.Printf("page:%d json Parse Error\n", page)
		return comics
	}
	return comics

}
