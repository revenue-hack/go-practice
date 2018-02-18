package main

import (
	"fmt"
	"log"
	"os"

	"flag"

	"gopl.io/ch5/links"
)

type depthList struct {
	depth int
	list  []string
}

var depth int

func main() {
	flag.IntVar(&depth, "depth", 3, "depth number")
	flag.Parse()
	var n int
	worklist := make(chan *depthList)
	go func() { worklist <- &depthList{0, os.Args[2:]} }()
	n++
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list.list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(depthOfList int, link string) {
					worklist <- crawl(depthOfList, link)
				}(list.depth, link)
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(depNum int, url string) *depthList {
	if depth < depNum {
		return nil
	}
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return &depthList{depNum + 1, list}
}
