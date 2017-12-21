package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"io/ioutil"
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
	url = "https://xkcd.com/%s/info.0.json"
	maxPage = 1930
	fileDir = "./src/ch04/ex12/data"
)



func main() {
	for i := 0; i < 10; i++ {
		createCacheFile(i + 1, downloadComics(i + 1))
	}
	//fmt.Printf("%v\n", readCacheFile(571).Month)

}

func createCacheFile(page int, comics Comics) {
	if _, err := os.Stat(fileDir); err != nil {
		os.Mkdir(fileDir, 0777)
	}
	filePath := fmt.Sprintf("%s/%d.json", fileDir, page)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("file create Error in createFileCache")
		panic(err)
	}
	defer file.Close()
	bytes, err := json.Marshal(comics)
	if err != nil {
		fmt.Println("json parse Error in createFileCache")
		panic(err)
	}
	file.Write(([]byte) (bytes))
}

func readCacheFile(page int) Comics {
	filePath := fmt.Sprintf("%s/%d.json", fileDir, page)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("このコミックは存在しません")
		panic(err)
	}
	var comics Comics
	if err := json.Unmarshal(file, &comics); err != nil {
		fmt.Println("json parse Error in readCacheFile")
		panic(err)
	}
	return comics
}


func downloadComics(page int) Comics {
	url := fmt.Sprintf(url, strconv.Itoa(page))
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
	return comics

}
