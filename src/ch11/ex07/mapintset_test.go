package ex02

import "testing"

var v, w MapIntSet

func TestMapIntSet(t *testing.T) {
	v.Add(1)
	v.Add(2)
	v.Add(100)
	v.Add(1000)
	v.Add(2000)
	w.Add(111)
	for _, c := range []struct {
		i        int
		expected bool
	}{
		{1, true},
		{2, true},
		{100, true},
		{1000, true},
		{2000, true},
		{234, false},
		{0, false},
		{12, false},
	} {
		if v.Has(c.i) != c.expected {
			t.Errorf("%d\t%v\n", c.i, c.expected)
		}
	}
	v.UnionWith(&w)
	if !v.Has(111) {
		t.Errorf("[UnionWith] v.is 111 nothing")
	}
	expected := "{1 2 100 111 1000 2000}"
	if v.String() != expected {
		t.Errorf("[String] v.is %s\twant is %s\n", v.String(), expected)
	}
}
