package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	element, err := build(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", element.string())
}

func build(reader io.Reader) (*Element, error) {
	dec := xml.NewDecoder(reader)
	var stack []*Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			element := &Element{Type: tok.Name, Attr: tok.Attr}
			if len(stack) != 0 {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, element)
			}
			stack = append(stack, element)
		case xml.EndElement:
			switch {
			case len(stack) == 1:
				return stack[0], nil
			case len(stack) > 1:
				stack = stack[:len(stack)-1]
			}
		case xml.CharData:
			if len(stack) != 0 {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, CharData(tok))
			}
		}
	}
	return nil, fmt.Errorf("build error")
}

func (e *Element) string() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("<%s", e.Type.Local))
	for _, attr := range e.Attr {
		buf.WriteString(fmt.Sprintf("%s='%s'", attr.Name, attr.Value))
	}
	if len(e.Children) == 0 {
		buf.WriteString("/>\n")
		return buf.String()
	}
	buf.WriteString(">\n")
	for _, child := range e.Children {
		switch t := child.(type) {
		case *Element:
			buf.WriteString(t.string())
		case CharData:
			buf.WriteString(string(t))
		}
	}
	buf.WriteString(fmt.Sprintf("</%s>\n", e.Type.Local))
	return buf.String()
}
