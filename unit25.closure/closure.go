package main

import "fmt"

func calc() func(x int) int {
	a, b := 3, 5 // 지역 변수는 함수가 끝나면 소멸되지만
	return func(x int) int {
		return a*x + b // 클로저이므로 함수를 호출 할 때마다 변수 a와 b의 값을 사용할 수 있음
	}
}

func main() {
	sum := func(a, b int) int {
		return a + b
	}

	r := sum(1, 2)
	fmt.Println(r)

	// 바깥 변수 사용
	a, b := 3, 5

	f := func(x int) int {
		return a*x + b
	}

	y := f(5)
	fmt.Println(y)
	fmt.Println()

	// 클로저를 변수에 저장
	f2 := calc()

	fmt.Println(f2(1))
	fmt.Println(f2(2))
	fmt.Println(f2(3))
	fmt.Println(f2(4))
	fmt.Println(f2(5))
}
