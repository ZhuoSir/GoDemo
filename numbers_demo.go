package main

import "fmt"

func main() {

	var balance = [5]float32{100.0, 21, 3.4, 19, 4.8}

	var balance2 = [...]float32{100.0, 21, 3.4, 19, 4.8}

	balance2[4] = 6

	fmt.Println(balance)
	fmt.Println(balance2)

}
