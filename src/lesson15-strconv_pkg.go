package main

import (
	"fmt"
	"strconv"
)

func lesson15() {
	var ori_str string = "666"
	var int_1 int

	fmt.Println(strconv.IntSize)
	int_1, _ = strconv.Atoi(ori_str)
	fmt.Println(int_1)
	int_1 += 5
	var new_str string = strconv.Itoa(int_1)
	fmt.Println(new_str)

	float_1, _ := strconv.ParseFloat(ori_str, 64)
	fmt.Println(float_1)
	float_1 += 8
	new_str = strconv.FormatFloat(float_1, 'f', 2, 64)
	fmt.Println(new_str)
}
