package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct{}

type Iphone struct{}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (iphone Iphone) call() {
	fmt.Println("I am iphone, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(Iphone)
	phone.call()
}
