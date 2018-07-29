package main

import "fmt"

func main() {

	var a int
	var ptr *int
	var pptr **int

	a = 3000
	ptr = &a
	pptr = &ptr

	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

	var x1, x2 = 10, 20
	fmt.Printf("交换前 a 的值 : %d\n", x1)
	fmt.Printf("交换前 b 的值 : %d\n", x2)

	swapint(&x1, &x2)
	fmt.Printf("交换后 a 的值 : %d\n", x1)
	fmt.Printf("交换后 b 的值 : %d\n", x2)
}

func swapint(x *int, y *int) {

	*x, *y = *y, *x

	//var temp int
	//temp = *x    /* 保存 x 地址的值 */
	//*x = *y      /* 将 y 赋值给 x */
	//*y = temp    /* 将 temp 赋值给 y */
}
