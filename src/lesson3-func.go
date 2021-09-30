package main

import "fmt"

func lesson3() {
	a, b := 100, 200
	result := max(a, b)
	fmt.Println(result)
}

func max(num_a int, num_b int) int {
	num_x := 0
	if num_a >= num_b {
		num_x = num_a
	} else {
		num_x = num_b
	}
	return num_x
}
