package main

import (
	"testing"
	"os"
)

func TestFetchOfFileExist(t *testing.T) {
	fileName, n, err := fetch("https://golang.org")
	if n == 0 || fileName == "" || err != nil {
		t.Error("file fetch error in TestFetchOfFileExist")
	}
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		t.Error("file open error in TestFetchOfFileExist")
	}
	defer file.Close()
}

func TestFetchOfFileNotExist(t *testing.T) {
	// not exist url
	fileName, n, err := fetch("https://example.ss/aaaaa")
	if n != 0 || fileName != "" || err == nil {
		t.Error("file fetch error in TestFetchOfFileNotExist")
	}
}

