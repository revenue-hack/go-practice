package main

import "fmt"

func main() {
	fmt.Printf("maxOne: %d\tminOne: %d\n", maxOne(1, 2, 3, 4, 5, -6), minOne(-1, -44, 12, 444, 999))
	max, err := max(1, 2, 3, 4, 5, -6)
	if err != nil {
		fmt.Println("error")
	}
	min, err := min(-1, -44, 12, 444, 999)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Printf("max: %d\tmin: %d\n", max, min)
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("vals nothing error")
	}

	var max int
	for i, val := range vals {
		if i == 0 || max < val {
			max = val
		}
	}
	return max, nil
}

func maxOne(o int, vals ...int) int {
	max := o
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("vals nothing error")
	}
	var min int
	for i, val := range vals {
		if i == 0 || min > val {
			min = val
		}
	}
	return min, nil
}

func minOne(o int, vals ...int) int {
	min := o
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min
}
