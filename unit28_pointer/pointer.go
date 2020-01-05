package main

import "fmt"

func hello(n *int) {
	*n = 2
}

func main() {
	var numPtr *int = new(int)
	fmt.Println(numPtr)
	*numPtr = 1
	fmt.Println(*numPtr)

	var num int = 3
	numPtr = &num

	fmt.Println(numPtr)
	fmt.Println(&num)

	var n int = 1
	hello(&n)
	fmt.Println(n)
}
