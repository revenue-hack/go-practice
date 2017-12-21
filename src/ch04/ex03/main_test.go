package main

import "testing"

func TestReverseOfNormal(t *testing.T) {
	a := [...]int{4, 5, 56, 6, 7, 8, 8}
	r := [...]int{8, 8, 7, 6, 56, 5, 4}
	reverse(&a)
	if a != r {
		t.Error("Test Normal Error")
	}
}

func TestReverseOfNormal2(t *testing.T) {
	a := [...]int{10, 5, 56, 6, 7, 8, 8}
	r := [...]int{8, 8, 7, 6, 56, 5, 4}
	reverse(&a)
	if a == r {
		t.Error("Test Normal2 Error")
	}
}
