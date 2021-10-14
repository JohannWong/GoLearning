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

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func lesson10() {
	a := functionValue(2, 1, add)
	b := functionValue(2, 1, sub)
	c := functionValue(2, 1, func(i1, i2 int) int {
		return i1 * i2
	})
	fmt.Println(a, b, c)

	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}
