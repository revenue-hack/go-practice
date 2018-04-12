package main

import (
	"log"

	"github.com/revenue-hack/go-practice/src/ch10/ex02/unarchive"
	_ "github.com/revenue-hack/go-practice/src/ch10/ex02/unarchive/tar"
	_ "github.com/revenue-hack/go-practice/src/ch10/ex02/unarchive/zip"
)

func main() {
	fileName := "sample.tar"

	r, err := unarchive.OpenReader(fileName)
	if err != nil {
		log.Fatal(err)
	}
	files := r.ReadFiles()
	log.Println("tar file info")
	for _, f := range files {
		log.Printf("%v\t%v\t%v\n", f.FileInfo().Mode().String(), f.FileInfo().Size(), f.Name())
	}

	fileName = "sample.zip"
	r, err = unarchive.OpenReader(fileName)
	if err != nil {
		log.Fatal(err)
	}
	files = r.ReadFiles()
	log.Println("zip file info")
	for _, f := range files {
		log.Printf("%v\t%v\t%v\n", f.FileInfo().Mode().String(), f.FileInfo().Size(), f.Name())
	}
}
