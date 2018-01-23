package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

func main() {
	parser(`<html>
	<a href="https://yahoo.co.jp"/>
	<a href="https://golang.org" />
	</html>`)
}

func parser(doms string) {
	doc, err := html.Parse(newReader(doms))
	if err != nil {
		panic(err)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

type reader struct {
	bytes []byte
	next  int
}

/**
 * rewrite io.ReaderのRead
 */
func (r *reader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	if r.next >= len(r.bytes) {
		return 0, io.EOF
	}
	nBytes := len(r.bytes) - r.next
	if nBytes > len(p) {
		nBytes = len(p)
	}
	copy(p, r.bytes[r.next:r.next+nBytes])
	r.next += nBytes
	return nBytes, nil
}

func newReader(s string) io.Reader {
	return &reader{[]byte(s), 0}
}
