package main

import "fmt"

func lesson4() {
	var num_list [100]int = [100]int{1: 5, 3: 10}
	fmt.Println(num_list)
	for i := 0; i < 100; i++ {
		num_list[i] = i
	}
	fmt.Println(num_list)

	var num_list_s = make([]int, 100)
	fmt.Println(num_list_s)
	num_list_s = append(num_list_s, 100)
	fmt.Println(num_list_s)
	for i := 0; i < 100; i++ {
		num_list_s[i] = i
	}
	fmt.Println(num_list_s)
	fmt.Println(len(num_list_s), cap(num_list_s))

	var num_list_2 = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(num_list_2)
	fmt.Println(num_list_2[1:4])
	fmt.Println(num_list_2[3:])
	fmt.Println(num_list_2[:3])

	num_list_3 := make([]int, len(num_list_2), (cap(num_list_2))*2)
	copy(num_list_3, num_list_s)
	fmt.Println(num_list_3)

	num_list_count, num_list_sum := 0, 0
	for _, num := range num_list_2 { //由于不使用索引，用空白符'_'省略
		num_list_count++
		num_list_sum += num
	}
	fmt.Println(num_list_count, num_list_sum)

	for i, num := range num_list_3 {
		fmt.Println(i, num)
	}
}
