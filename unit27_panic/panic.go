package main

import "fmt"

func f() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	a := [...]int{1, 2}
	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
	// panic("Error !!!")
}
func main() {
	f()
	fmt.Println("Hello, world")
	// a := [...]int{1, 2}

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(a[i])
	// }

	// panic("Error !!")
	// fmt.Println("Hello, world!") // 실행되지 않음
}
