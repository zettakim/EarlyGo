package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func hello() {
// 	fmt.Println("Hello, world!")
// }

func hello(n int) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println(n)
}

func main() {
	// go hello()
	// fmt.Scanln()

	for i := 0; i < 100; i++ {
		go hello(i)
	}

	fmt.Scanln()
}
