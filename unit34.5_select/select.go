package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)

	// go func() {
	// 	for {
	// 		c1 <- 10
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }()
	go func() {
		for {
			i := <-c1
			fmt.Println("c1 : ", i)
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Hello, world!"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			// select {
			// case i := <-c1:
			// 	fmt.Println("c1 :", i)
			// case s := <-c2:
			// 	fmt.Println("c2 :", s)
			// case <-time.After(50 * time.Millisecond):
			// 	fmt.Println("timeout")
			// }
			select {
			case c1 <- 10:
			case s := <-c2:
				fmt.Println("c2 :", s)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
