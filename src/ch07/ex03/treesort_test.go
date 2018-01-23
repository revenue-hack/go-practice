package main

import (
	"testing"
)

var cases = []struct {
	expected []int
	values []int
}{
	{[]int{2,3,4,5,6}, []int{6,2,5,3,4}},
	{[]int{-111,2,6,10,500}, []int{2,-111,10,500,6}},
	{nil, nil},
}

func TestSort(t *testing.T) {
	for _, c := range cases {
		vals := c.values
		Sort(vals)
		for i, v := range vals {
			if v != c.expected[i] {
				t.Errorf("don't match expected value sort value: %d\texpected value: %d\n", v, c.expected[i])
			}
		}
	}

}
