package main

import "testing"

func TestIntSet_AddAll(t *testing.T) {
	var array []int
	var x, y IntSet
	for i := 0; i < 1000; i++ {
		array = append(array, i)
		x.Add(i)
	}
	y.AddAll(array...)
	for i, xw := range x.words {
		if xw != y.words[i] {
			t.Errorf("addall don't match %d\n", xw)
		}
	}
}
