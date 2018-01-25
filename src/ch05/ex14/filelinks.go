package main

import (
	"fmt"
	"os"
	"log"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func listFiles(path string) []string {
	fmt.Println(path)
	dirInfos := extractFileInfos(path)

	var files []string
	for _, dirInfo := range dirInfos {
		name := dirInfo.Name()
		if name[0] == '.' {
			continue
		}
		files = append(files, path+"/"+dirInfo.Name())
	}
	return files
}

func extractFileInfos(path string) []os.FileInfo {
	f, err := os.Open(path)
	if err != nil {
		log.Print(err)
		return nil
	}

	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		log.Print(err)
		return nil
	}

	if !fileInfo.IsDir() {
		return nil
	}

	dirInfos, err := f.Readdir(0)
	if err != nil {
		log.Print(err)
		return nil
	}
	return dirInfos
}

func main() {
	breadthFirst(listFiles, os.Args[1:])
}