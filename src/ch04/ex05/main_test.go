package main

import "testing"

func TestRemoveOfNormal(t *testing.T) {
	strs := []string{"a", "a", "ss", "2"}
	result:= []string{"a", "ss", "2"}
	for i, str := range remove(strs, 1) {
		if str != result[i] {
			t.Error("Test Remove Normal Error")
		}
	}
}

func TestUniqueOfNormal(t *testing.T) {
	strs := []string{"a", "a", "ss", "2"}
	result:= []string{"a", "ss", "2"}
	for i, str := range unique(strs) {
		if str != result[i] {
			t.Error("Test Unique Normal Error")
		}
	}
}
