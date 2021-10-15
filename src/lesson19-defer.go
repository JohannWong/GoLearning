package main

import (
	"fmt"
	"log"
)

func lesson19() {
	log.SetFlags(log.Lshortfile)
	doDBOperations()

	to_trace_b()

	fmt.Println(defer_func_test())

	log.Print()
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
	log.Print()
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
	log.Print()
}

func to_trace_b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	to_trace_a()
	log.Print()
}

func defer_func_test() (res int) {
	defer func() {
		res++
	}()
	return 1
}
