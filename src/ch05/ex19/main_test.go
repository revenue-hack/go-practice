package main

import "testing"

func TestVariable(t *testing.T) {
	if variable(10) != 10 {
		t.Error("TestVariable not equal Error")
	}
	if variable(10) == 8 {
		t.Error("TestVariable equal Error")
	}

}
