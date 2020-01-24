package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	s := "Hello, world"

	for i := 0; i < 100; i++ {
		// go func(n int) {
		// 	fmt.Println(s, n)
		// }(i)
		go func() {
			fmt.Println(s, i)
		}()
	}

	fmt.Scanln()
}
