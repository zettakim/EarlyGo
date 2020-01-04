package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func HelloWorld() {
	defer func() {
		fmt.Println("world2")
	}()
	func() {
		fmt.Println("Hello2")
	}()
}

func main() {
	defer world() // 현재 한수 끝나기 직전 호출
	hello()
	hello()
	hello()

	HelloWorld()

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
