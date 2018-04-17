package main

import (
	"net/http"
	"net/url"
	"testing"

	"reflect"

	"github.com/revenue-hack/go-practice/src/ch12/ex12/params"
)

func TestUnPack(t *testing.T) {
	for _, c := range []struct {
		url        string
		expected   Data
		isValidate bool
	}{
		{
			"http://hoge.com?l=hello&l=world&max=10&x=true",
			Data{[]string{"hello", "world"}, 10, true},
			false,
		},
		{
			"http://hoge.com?l=world&max=1111&x=false",
			Data{[]string{"world"}, 1111, false},
			false,
		},
	} {
		var req http.Request
		url, err := url.Parse(c.url)
		if err != nil {
			t.Errorf("parse error %v\n", err)
		}
		req.URL = url
		var data Data

		err = params.Unpack(&req, &data)
		if isValidate := err != nil; isValidate != c.isValidate {
			t.Errorf("unpack error %v\n", err)
		}
		if !reflect.DeepEqual(data, c.expected) {
			t.Errorf("result = %v expected = %v\n", data, c.expected)
		}
	}

	for _, c := range []struct {
		input    NoTagData
		url      string
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
		result, err := Pack(&c.input, c.url)
		if err != nil {
			t.Error(err)
		}
		if result != c.expected {
			t.Errorf("result = %v, expected = %v\n", result, c.expected)
		}
	}
}

func TestExUnPack(t *testing.T) {
	for _, c := range []struct {
		url        string
		expected   ExData
		isValidate bool
	}{
		{
			"http://hoge.com?mail=hgoe@c.com&credit=1234567890&zip=1234567",
			ExData{"hgoe@c.com", "1234567890", "1234567"},
			false,
		},
		{
			"http://hoge.com?mail=hgoec.com&credit=1234567890&zip=1234567",
			ExData{},
			true,
		},
		{
			"http://hoge.com?mail=hgoe@c.com&credit=00000001234567890&zip=1234567",
			ExData{},
			true,
		},
		{
			"http://hoge.com?mail=hgoe@c.com&credit=1234567890&zip=123456ahofd7",
			ExData{},
			true,
		},
	} {
		var req http.Request
		url, err := url.Parse(c.url)
		if err != nil {
			t.Errorf("parse error %v\n", err)
		}
		req.URL = url
		var data ExData
		err = params.Unpack(&req, &data)
		if isValidate := err != nil; isValidate != c.isValidate {
			t.Errorf("unpack error %v\n", err)
		}
		if !c.isValidate && !reflect.DeepEqual(data, c.expected) {
			t.Errorf("result = %v expected = %v\n", data, c.expected)
		}
	}
}
