package main

import "fmt"

func main() {
	i := 0
	for { // 무한 루프
		if i > 4 {
			break
		}

		fmt.Println(i)
		i++
	}
Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("hello, world!")
}
