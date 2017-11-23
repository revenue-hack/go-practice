package main

import (
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	os.Args = []string{"param1", "param2"}
	if echo() != "param1 param2" {
		t.Error("TestEcho test fail")
	}
}

func TestNonEcho(t *testing.T) {
	os.Args = []string{"ppp"}
	if echo() == "param1" {
		t.Error("TestNonEcho test fail")
	}
}

func TestEmptyEcho(t *testing.T) {
	os.Args = []string{""}
	if echo() == "param1" {
		t.Error("TestEmptyEcho test fail")
	}
	if echo() != "" {
		t.Error("TestEmptyEcho test fail")
	}
}

func TestNilEcho(t *testing.T) {
	if echo() != "" {
		t.Error("TestNilEcho test fail")
	}
}
