package main

import "fmt"

func main() {
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
}

func sum(a int, b int) int {
	return a + b
}

func sum2(a int, b int) (r int) {
	r = a + b
	return
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
