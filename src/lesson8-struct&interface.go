package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("Its NokiaPhone")
}

type IPhone struct {
}

func (iphone IPhone) call() {
	fmt.Println("Its IPhone")
}

func lesson8() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
