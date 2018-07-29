package main

import (
	"fmt"
	"unsafe"
)

var d, f int

var (
	x int
	y int
	s = "go"
)

const (
	c1 = "abc"
	c2 = len(c1)
	c3 = unsafe.Sizeof(c1)
)

/*iota，特殊常量，可以认为是一个可以被编译器修改的常量。

在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。

iota 可以被用作枚举值*/
const (
	i1 = iota
	i2
	i3
)

const j = "its a constant"

func main() {
	var a int = 10
	var b = 10
	c := 10

	fmt.Println(a, b, c)
	fmt.Println(d, f, x, y, s)
	fmt.Println(j)

	fmt.Println(c1, c2, c3)
	fmt.Println(i1, i2, i3)

}
