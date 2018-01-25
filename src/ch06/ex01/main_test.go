package main

import (
	"testing"
)

var x IntSet

func TestIntSet_Has(t *testing.T) {
	x.Add(1)
	x.Add(2)
	x.Add(100)
	x.Add(1000)
	x.Add(2000)
	x.Add(500000)
	for _, c := range []struct {
		i        int
		expected bool
	}{
		{1, true},
		{2, true},
		{100, true},
		{1000, true},
		{2000, true},
		{500000, true},
		{234, false},
		{0, false},
		{12, false},
	} {
		if x.Has(c.i) != c.expected {
			t.Errorf("Has: %d\t%v\n", c.i, c.expected)
		}
	}
}

func TestIntSet_Len(t *testing.T) {
	for _, c := range []struct {
		values   []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 0, 999, 22}, 4},
	} {
		var x IntSet
		for _, v := range c.values {
			x.Add(v)
		}
		if x.Len() != c.expected {
			t.Errorf("Len: %v\t%d\n", c.values, c.expected)
		}
	}
}

func TestIntSet_Remove(t *testing.T) {
	for _, c := range []struct {
		values    []int
		deleteInt int
	}{
		{[]int{1}, 1},
		{[]int{1, 144}, 1},
		{[]int{1, 0, 999, 22}, 22},
	} {
		var x IntSet
		for _, v := range c.values {
			x.Add(v)
		}
		x.Remove(c.deleteInt)
		if x.Has(c.deleteInt) {
			t.Errorf("Remove: %v\t%d\n", c.values, c.deleteInt)

		}
	}
}

func TestIntSet_Clear(t *testing.T) {
	var x IntSet
	for i := 0; i < 1000; i++ {
		x.Add(i)
	}
	x.Clear()
	if x.Len() != 0 {
		t.Errorf("Clear: %v\n", x)
	}
}

func TestIntSet_Copy(t *testing.T) {
	var x IntSet
	for i := 0; i < 1000; i++ {
		x.Add(i)
	}
	y := x.Copy()
	for i, xw := range x.words {
		if xw != y.words[i] {
			t.Errorf("copy: %v\n", xw)
		}

	}

}
