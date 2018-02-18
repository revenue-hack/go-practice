package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verebose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	done := make(chan struct{})
	for _, root := range roots {
		go walk(root, done)
	}
	for i := 0; i < len(roots); i++ {
		<-done
	}
}

var nfiles, nbytes int64

func walk(root string, done chan<- struct{}) {
	defer func() { done <- struct{}{} }()
	var n sync.WaitGroup
	fileSizes := make(chan int64)
	n.Add(1)
	go walkDir(root, &n, fileSizes)
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	var tick <-chan time.Time
	if *verebose {
		tick = time.Tick(500 * time.Millisecond)
	}
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(root, nfiles, nbytes)
		}
	}
	printDiskUsage(root, nfiles, nbytes)
}

func printDiskUsage(dir string, nfiles, nbytes int64) {
	fmt.Printf("%s: %d files %.1f GB\n", dir, nfiles, float64(nbytes)/1e9)

}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
