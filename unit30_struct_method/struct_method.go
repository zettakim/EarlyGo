// 리시버
package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

// 리시버 변수 정의 (연결할 구조체 정의)
func (rect *Rectangle) area() int {
	return rect.width * rect.height
}

func (rect *Rectangle) sacleA(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func (rect Rectangle) sacleB(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func (_ Rectangle) dummy() {
	fmt.Println("dummy")
}

func main() {
	rect := Rectangle{10, 20}
	fmt.Println(rect.area())

	rect1 := Rectangle{30, 30}
	rect1.sacleA(10)
	fmt.Println(rect1)

	rect2 := Rectangle{30, 30}
	rect2.sacleB(10)
	fmt.Println(rect2)

	var r Rectangle
	r.dummy()
}
