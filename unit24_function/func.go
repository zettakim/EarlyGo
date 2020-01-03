package main

import "fmt"

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func sum2(a int, b int) (r int) {
	r = a + b
	return
}

func sum3(n ...int) int { //가변인자
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}

func sumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}

func sumAndDiff2(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func hello() (int, int, int, int, int) {
	return 1, 2, 3, 4, 5
}

func main() {
	//슬라이스 사용
	f := []func(int, int) int{sum, diff}

	fmt.Println(f[0](1, 2))
	fmt.Println(f[1](1, 2))
	fmt.Println()

	// 맵 사용
	f2 := map[string]func(int, int) int{
		"sum":  sum,
		"diff": diff,
	}
	fmt.Println(f2["sum"](1, 2))
	fmt.Println(f2["diff"](1, 2))
	fmt.Println()

	//함수를 변수에 저장
	var hello2 func(a int, b int) int = sum
	world2 := sum

	fmt.Println("hello2(1, 2)->", hello2(1, 2))
	fmt.Println("world2(1, 2)->", world2(1, 2))

	r := sum(1, 5)
	fmt.Println(r)
	r = sum(3, 4)
	fmt.Println(r)
	sum, diff := sumAndDiff(7, 2)
	fmt.Println(sum, diff)
	_, diff = sumAndDiff(9, 3)
	fmt.Println(diff)
	a, _, c, _, e := hello()
	fmt.Println(a, c, e)

	sum, diff = sumAndDiff2(6, 2)
	fmt.Println(sum, diff)

	r = sum3(1, 2, 3, 4, 5)
	fmt.Println(r)

	n := []int{1, 2, 3, 4, 5}
	r = sum3(n...)
	fmt.Println(r)

	fmt.Println("factorial(5)=", factorial(5))

	// 익명 함수
	func() {
		fmt.Println("Hello")
	}()

	func(s string) {
		fmt.Println(s)
	}("world")

	ret := func(a int, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(ret)
}
