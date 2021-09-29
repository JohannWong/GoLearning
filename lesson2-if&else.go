package main

import "fmt"

func lesson2() {
	num_a, num_b := 10, 10
	if num_a > num_b {
		fmt.Println("输入的数字小于10")
	} else if num_a < num_b {
		fmt.Println("输入的数字大于10")
	} else {
		fmt.Println("输入的数字等于10")
	}

	var addr_a *int = &num_a
	fmt.Println(addr_a)
}
