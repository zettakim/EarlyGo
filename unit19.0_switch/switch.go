package main

import "fmt"

func main() {
	s := "Hello"
	i := 2

	switch i {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		if s == "Hello" {
			fmt.Println("Hello 2")
			break
		}
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	default:
		fmt.Println(-1)
	}
}
