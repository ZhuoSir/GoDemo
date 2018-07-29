package main

import (
	"fmt"
	"math"
)

func main() {

	var a = 100
	var b = 200
	var ret int

	ret = max(a, b)
	fmt.Println(ret)

	s1, s2 := swap("hello", "world")
	fmt.Println(s1, s2)

	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	addfunc := add_func(1, 2)
	fmt.Println(addfunc())
	fmt.Println(addfunc())
	fmt.Println(addfunc())

	getSquarRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquarRoot(64))
}

func max(i int, i2 int) int {

	var result int
	if i > i2 {
		result = i
	} else {
		result = i2
	}

	return result
}

func swap(s1 string, s2 string) (string, string) {
	return s2, s1
}

//Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
//以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func add_func(x1 int, x2 int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, x1 + x2
	}
}
