package main

import (
	"testing"
)

func TestEcho(t *testing.T) {
	if echo(1, "param1") != "i: 1 value: param1\n" {
		t.Error("TestEcho test fail")
	}
}

func TestNotEcho(t *testing.T) {
	if echo(1, "param1") == "i: 0 value: ssssparam1" {
		t.Error("TestEcho test fail")
	}
}
