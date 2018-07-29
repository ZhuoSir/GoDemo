package main

import "fmt"

type Circle struct {
	radius float64
}

//Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，
// 接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {

	var c1 Circle
	c1.radius = 10.00

	fmt.Print("Area of Circle(c1) = ", c1.getArea())

}
