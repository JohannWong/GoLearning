package main

import "fmt"

var num int = 10
var numx2, numx3 int

func lesson3() {
	a, b := 100, 200
	result := max(a, b)
	fmt.Println(result)

	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()

	multi_n := 0
	multi_n_addr := &multi_n
	Multiply(10, 5, multi_n_addr)
	fmt.Println("Multiply:", *multi_n_addr, multi_n)
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

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}
