package main

import (
	"os"
	"testing"
)

func TestCountLines(t *testing.T) {
	fileName := "a.txt"
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	counts := make(map[string]int)
	countFiles := make(map[string]string)
	countLines(f, counts, countFiles, fileName)
	if len(counts) == 0 {
		t.Error("TestCountLines test counts array fail")
	}
	if len(countFiles) == 0 {
		t.Error("TestCountLines test countLines array fail")
	}
	if counts["9"] != 2 {
		t.Error("TestCountLines test counts detail fail")
	}
	if countFiles["1"] != "a.txt" {
		t.Error("TestCountLines test countLines detail fail")
	}
}

func TestCountLinesMulti(t *testing.T) {
	files := []string{"a.txt", "b.txt"}
	counts := make(map[string]int)
	countFiles := make(map[string]string)
	for _, args := range files {
		f, err := os.Open(args)
		if err != nil {
			panic(err)
		}
		countLines(f, counts, countFiles, args)
	}
	if len(counts) == 0 {
		t.Error("TestCountLines test counts array fail")
	}
	if len(countFiles) == 0 {
		t.Error("TestCountLines test countLines array fail")
	}
	if counts["1"] != 2 {
		t.Error("TestCountLines test counts detail fail")
	}
	if countFiles["1"] != "a.txt,b.txt" {
		t.Error("TestCountLines test countLines detail fail")
	}

}
