package main

import (
	"fmt"
	"math/rand"
	"time"
)

func switch1() {
	i := 3

	switch i {
	case 4:
		fmt.Println("4 이상")
		fallthrough
	case 3:
		fmt.Println("3 이상")
		fallthrough
	case 2:
		fmt.Println("2 이상")
		fallthrough
	case 1:
		fmt.Println("1 이상")
		fallthrough
	case 0:
		fmt.Println("0 이상")
	}
}

func switch2() {
	j := 3

	switch j {
	case 2, 4, 6:
		fmt.Println("짝수")
	case 1, 3, 5:
		fmt.Println("홀수")
	}
}

func switch3() {
	k := 7

	switch {
	case k >= 5 && k < 10:
		fmt.Println("5 이상, 10 미만")
	case k >= 0 && k < 5:
		fmt.Println("0 이상, 5 미만")
	}
}

func switch4() { // 조건식으로 분기하기
	rand.Seed(time.Now().UnixNano())

	switch i := rand.Intn(10); {
	case i >= 3 && i < 6:
		fmt.Println("3 이상, 6 미만")
	case i == 9:
		fmt.Println("9")
	default:
		fmt.Println(i)
	}
}

func main() {
	switch1()
	switch2()
	switch3()
	switch4()
}
