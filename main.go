package main

import "fmt"

const WIDTH, LENGTH int = 5, 10

func main() {
	var a = "test"
	var b, c int = 1, 2
	var d string = "nice"
	var i_point *int
	var i_list []int
	var f_value func(string)
	var err error
	flag := true
	area := LENGTH * WIDTH

	fmt.Println(a)
	fmt.Println(b, c)
	fmt.Println(d)
	fmt.Println(i_point)
	fmt.Println(i_list)
	fmt.Println(f_value)
	fmt.Println(err)
	fmt.Println(flag)
	fmt.Println(area)
}
