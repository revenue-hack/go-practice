package main

import "testing"

var maxCases = []struct {
	nums     []int
	expected int
	o        int
}{
	{[]int{1, 2, 3, 4, 5}, 5, 2},
	{[]int{1, 3333, 3, 4, 5}, 3333, 44},
	{nil, 1, 1},
}

var minCases = []struct {
	nums     []int
	expected int
	o        int
}{
	{[]int{1, 2, 3, 4, 5}, 1, 2},
	{[]int{1, 3333, 3, 4, 5}, 1, 44},
	{nil, 1, 1},
}

func TestMax(t *testing.T) {
	for _, c := range maxCases {
		result, err := max(c.nums...)
		if c.nums == nil && err == nil {
			t.Errorf("[isError] nums: %v\texpected: %d\t\n", c.nums, c.expected)
			return
		}
		if c.nums != nil && result != c.expected {
			t.Errorf("[expected] nums: %v\texpected: %d\t\n", c.nums, c.expected)
			return
		}
	}
}

func TestMaxOne(t *testing.T) {
	for _, c := range maxCases {
		result := maxOne(c.o, c.nums...)
		if result != c.expected {
			t.Errorf("[expected] nums: %v\texpected: %d\to: %d\n", c.nums, c.expected, c.o)
			return
		}
	}
}

func TestMin(t *testing.T) {
	for _, c := range minCases {
		result, err := min(c.nums...)
		if c.nums == nil && err == nil {
			t.Errorf("[isError] nums: %v\texpected: %d\t\n", c.nums, c.expected)
			return
		}
		if c.nums != nil && result != c.expected {
			t.Errorf("[expected] nums: %v\texpected: %d\t\n", c.nums, c.expected)
			return
		}
	}
}

func TestMinOne(t *testing.T) {
	for _, c := range minCases {
		result := minOne(c.o, c.nums...)
		if result != c.expected {
			t.Errorf("[expected] nums: %v\texpected: %d\t\n", c.nums, c.expected)
			return
		}
	}
}
