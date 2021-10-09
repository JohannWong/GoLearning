package main

import "fmt"

func functionValue(num1, num2 int, do func(int, int) int) int {
	return do(num1, num2)
}

func add(num1, num2 int) int {
	return num1 + num2
}

func sub(num1, num2 int) int {
	return num1 - num2
}

func lesson10() {
	a := functionValue(2, 1, add)
	b := functionValue(2, 1, sub)
	c := functionValue(2, 1, func(i1, i2 int) int {
		return i1 * i2
	})
	fmt.Println(a, b, c)
}
