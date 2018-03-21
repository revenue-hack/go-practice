package ex02

import "testing"

var x, y IntSet

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

func TestIntSet_IntersectWith(t *testing.T) {
	for _, c := range []struct {
		c1       []int
		c2       []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2}, []int{1, 2}},
		{[]int{}, []int{1, 2}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5}, []int{1, 5}},
	} {
		var x IntSet
		var y IntSet
		x.AddAll(c.c1...)
		y.AddAll(c.c2...)
		x.IntersectWith(&y)
		for _, e := range c.expected {
			if !x.Has(e) {
				t.Errorf("intersectWith don't match %d\t%v\n", e, x)
			}
		}
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	for _, c := range []struct {
		c1       []int
		c2       []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2}, []int{3, 4, 5}},
		{[]int{}, []int{1, 2}, []int{}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5}, []int{2, 3, 4}},
	} {
		var x IntSet
		var y IntSet
		x.AddAll(c.c1...)
		y.AddAll(c.c2...)
		x.DifferenceWith(&y)
		if x.Len() != len(c.expected) {
			t.Errorf("differenceWith don't match %v\n", x)
		}
	}
}

func TestIntSet_SymmetricDifference(t *testing.T) {
	for _, c := range []struct {
		c1       []int
		c2       []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2}, []int{3, 4, 5}},
		{[]int{}, []int{1, 2}, []int{1, 2}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},

		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}, []int{1, 2, 3, 6, 7, 8}},
	} {
		var x IntSet
		var y IntSet
		x.AddAll(c.c1...)
		y.AddAll(c.c2...)
		x.SymmetricDifference(&y)
		if x.Len() != len(c.expected) {
			t.Errorf("symmetricDifference Len don't match %v\n", x)
		}
		for i, e := range c.expected {
			if !x.Has(e) {
				t.Errorf("symmetricDifference Has don't match %d\t%v\n", i, x)
			}
		}
	}
}
