package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	tags := parseTag(os.Args[1:])
	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if result, isOk := containsAll(stack, tags); isOk && result != "" {
				fmt.Printf("%s: %s\n", result, tok)
			}
		}
	}
}

func containsAll(stack []xml.StartElement, tags []Tag) (string, bool) {
	if len(stack) == 0 {
		return "", true
	}
	var result string
	for len(stack) >= len(tags) {
		if len(tags) == 0 {
			return result, true
		}
		if tags[0].key == "tag" && stack[0].Name.Local == tags[0].value {
			tags = tags[1:]
		} else if stack[0].Attr != nil && containsSelector(stack[0].Attr, tags) {
			result += stack[0].Name.Local + " "
			if tags[0].key == "class" {
				result += "."
			} else {
				result += "#"
			}
			result += tags[0].value + " "
			tags = tags[1:]
		}
		result += stack[0].Name.Local + " "
		stack = stack[1:]
	}
	return "", false
}

func containsSelector(attrs []xml.Attr, tags []Tag) bool {
	for _, attr := range attrs {
		if attr.Name.Local == tags[0].key && attr.Value == tags[0].value {
			return true
		}
	}
	return false
}

type Tag struct {
	key   string
	value string
}

func parseTag(args []string) []Tag {
	var tags []Tag
	for _, arg := range args {
		// when class
		if strings.HasPrefix(arg, ".") {
			tags = append(tags, Tag{key: "class", value: arg[1:]})
			// when id
		} else if strings.HasPrefix(arg, "#") {
			tags = append(tags, Tag{key: "id", value: arg[1:]})
		} else {
			tags = append(tags, Tag{key: "tag", value: arg[0:]})
		}
	}
	return tags
}
