package main

import "fmt"

const MAX = 3

func main() {
	var a = 10
	fmt.Printf("变量的地址: %x\n", &a)

	var ip *int
	ip = &a
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	fmt.Printf("*ip 变量的值: %d\n", *ip)

	//var a = [MAX]int{10,100,200}
	//var i int
	//var ptr [MAX]*int
	//
	//for i = 0; i < MAX; i++ {
	//	ptr[i] = a[i]
	//}
	//
	//for  i = 0; i < MAX; i++ {
	//	fmt.Printf("a[%d] = %d\n", i,*ptr[i] )
	//}
}
