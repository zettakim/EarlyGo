package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(a int, b int, c chan int) {
	c <- a + b
}

func main() {
	c := make(chan int)

	go sum(1, 2, c)

	n := <-c
	fmt.Println(n)
	fmt.Println("-----------------")

	done := make(chan bool)
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("고루틴 : ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("메인 함수 : ", i)
	}
	fmt.Println("-----------------")

	runtime.GOMAXPROCS(1)

	done2 := make(chan bool, 2)
	count2 := 4

	go func() {
		for i := 0; i < count2; i++ {
			done2 <- true
			fmt.Println("고루틴2 : ", i)
		}
	}()

	for i := 0; i < count2; i++ {
		<-done2
		fmt.Println("메인 함수2 : ", i)
	}
	fmt.Println("-----------------")

	c2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c2 <- i
		}
		close(c2)
	}()

	for i := range c2 {
		fmt.Println(i)
	}
	fmt.Println("-----------------")

	c3 := make(chan int, 1)

	go func() {
		c3 <- 1
	}()

	a, ok := <-c3
	fmt.Println(a, ok)

	close(c3)
	a, ok = <-c3
	fmt.Println(a, ok)
}
