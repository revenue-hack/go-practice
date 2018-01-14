package main

import (
	"strings"
	"testing"
)

func TestExpandOfNormal(t *testing.T) {
	strs := map[string]string{
		"$hoge$sssss$aaaaa": "HOGESSSSSAAAAA",
		"$$$hoge$sssss":     "$$HOGESSSSS",
		"$HOGE$SS":          "HOGESS",
	}
	for in, out := range strs {
		if s := expand(in, func(s string) string {
			return strings.ToUpper(s)
		}); s != out {
			t.Errorf("TestExpandNormal Error\tin: %s\tout: %s\tresult:%s\n", in, out, s)
		}
	}
}

func TestExpandOfStringEmpty(t *testing.T) {
	if s := expand("", func(s string) string {
		return strings.ToUpper(s)
	}); s != "" {
		t.Errorf("TestExpandOfStringEmpty Error\tresult:%s\n", s)
	}
}

func TestExpandOfNothingDollar(t *testing.T) {
	strs := map[string]string{
		"hogesssssaaaaa": "hogesssssaaaaa",
		"hogesssss":      "hogesssss",
		"HOGESS":         "HOGESS",
	}
	for in, out := range strs {
		if s := expand(in, func(s string) string {
			return strings.ToUpper(s)
		}); s != out {
			t.Errorf("TestExpandOfNothingDollar Error\tin: %s\tout: %s\tresult:%s\n", in, out, s)
		}
	}
}
