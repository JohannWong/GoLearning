package main

import "fmt"

const FIBO_LIM = 41

var fib_list [FIBO_LIM]int

func fibonacci(n int) (res int) {
	if fib_list[n] != 0 {
		res = fib_list[n]
		return
	} else if n < 2 {
		return n
	} else {
		res = fibonacci(n-2) + fibonacci(n-1)
	}
	fib_list[n] = res
	return
}

func lesson6() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}
