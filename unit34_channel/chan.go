package main

import (
	"fmt"
	"runtime"
	"time"
)

func num(a, b int) <-chan int {
	out := make(chan int)
	go func() {
		out <- a
		out <- b
		close(out)
	}()
	return out
}

func sum3(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		r := 0
		for i := range c {
			r = r + i
		}
		out <- r
	}()
	return out
}

func sum(a int, b int, c chan int) {
	c <- a + b
}

func sum2(a, b int) <-chan int {
	out := make(chan int)
	go func() {
		out <- a + b
	}()
	return out
}

func producer(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	c <- 100
}

func consumer(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(<-c)
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
	fmt.Println("-----------------")

	c4 := make(chan int)

	go producer(c4)
	go consumer(c4)

	fmt.Println("-----------------")
	c5 := sum2(1, 2)
	fmt.Println("sum2:", <-c5)
	fmt.Println("-----------------")

	c6 := num(1, 2)
	out6 := sum3(c6)

	fmt.Println("sum3:", <-out6)
	fmt.Scanln()

}
