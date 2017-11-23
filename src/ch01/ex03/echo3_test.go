package main

import (
	"os"
	"testing"
)

func TestInefficiencyEcho(t *testing.T) {
	if inefficiencyEcho("param1") != "param1 " {
		t.Error("TestInefficiencyEcho test fail")
	}
}

func TestNotInefficiencyEcho(t *testing.T) {
	if inefficiencyEcho("param2") == "param1 " {
		t.Error("TestInefficiencyEcho test fail")
	}
}
func TestEfficiencyEcho(t *testing.T) {
	os.Args = []string{"param1", "param2"}
	if efficiencyEcho() != "param2" {
		t.Error("TestEfficiencyEcho test fail")
	}
}

func TestNotEfficiencyEcho(t *testing.T) {
	os.Args = []string{"param1", "param2"}
	if efficiencyEcho() == "param1" {
		t.Error("TestEfficiencyEcho test fail")
	}
}
