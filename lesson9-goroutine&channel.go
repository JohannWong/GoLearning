package main

import "fmt"

func sum(int_list []int, channel chan int) {
	sum := 0
	for _, value := range int_list {
		sum += value
	}
	channel <- sum
}

func fibonacci_go(max_i int, channel chan int) {
	x, y := 0, 1
	for i := 0; i < max_i; i++ {
		channel <- x
		x, y = y, x+y
	}
	close(channel)
}

func lesson9() {
	list := []int{1, 10, 100, 1000, 10000}
	channel_1 := make(chan int)
	go sum(list, channel_1)
	go sum(list[:3], channel_1)
	go sum(list[3:], channel_1)
	x, y, z := <-channel_1, <-channel_1, <-channel_1

	fmt.Println(x, y, z)
	close(channel_1)

	channel_2 := make(chan int, 10)
	go fibonacci_go(cap(channel_2), channel_2)
	for fibo_num := range channel_2 {
		fmt.Println(fibo_num)
	}
}
