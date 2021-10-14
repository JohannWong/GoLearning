package main

import "fmt"

func lesson19() {
	doDBOperations()

	to_trace_b()

	fmt.Println(defer_func_test())
}

func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func to_trace_a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func to_trace_b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	to_trace_a()
}

func defer_func_test() (res int) {
	defer func() {
		res++
	}()
	return 1
}
