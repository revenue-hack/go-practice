package ex02

import "testing"

var x, y IntSet

func TestIntSet(t *testing.T) {
	x.Add(1)
	x.Add(2)
	x.Add(100)
	x.Add(1000)
	x.Add(2000)
	x.Add(500000)
	y.Add(111)
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
			t.Errorf("%d\t%v\n", c.i, c.expected)
		}
	}
	x.UnionWith(&y)
	if !x.Has(111) {
		t.Errorf("[UnionWith] x is 111 nothing")
	}
	expected := "{1 2 100 111 1000 2000 500000}"
	if x.String() != expected {
		t.Errorf("[String] x is %s\twant is %s\n", x.String(), expected)
	}
}
