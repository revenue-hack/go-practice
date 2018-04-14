package main

import (
	"testing"
)

func TestPack(t *testing.T) {
	for _, c := range []struct {
		input    Data
		domain   string
		expected string
	}{
		{
			Data{[]string{"hello", "world"}, 10, true},
			"http://hoge.com",
			"http://hoge.com?l=hello&l=world&max=10&x=true",
		},
		{
			Data{[]string{"world"}, 1111, false},
			"http://aaaa.com",
			"http://aaaa.com?l=world&max=1111&x=false",
		},
	} {
		result, err := Pack(&c.input, c.domain)
		if err != nil {
			t.Error(err)
		}
		if result != c.expected {
			t.Errorf("result = %v, expected = %v\n", result, c.expected)
		}
	}

	for _, c := range []struct {
		input    NoTagData
		domain   string
		expected string
	}{
		{
			NoTagData{[]string{"hello", "world"}, 10, true},
			"http://hoge.com",
			"http://hoge.com?labels=hello&labels=world&maxresults=10&exact=true",
		},
		{
			NoTagData{[]string{"world"}, 1111, false},
			"http://aaaa.com",
			"http://aaaa.com?labels=world&maxresults=1111&exact=false",
		},
	} {
		result, err := Pack(&c.input, c.domain)
		if err != nil {
			t.Error(err)
		}
		if result != c.expected {
			t.Errorf("result = %v, expected = %v\n", result, c.expected)
		}
	}
}
