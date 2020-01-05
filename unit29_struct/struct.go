package main

import "fmt"

// Rectangle strcut
type Rectangle struct {
	width  int
	height int
}

// NewRectangle strcut
func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}
}

func rectangleArea(rect *Rectangle) int {
	return rect.width * rect.height
}

func rectangleAreaA(rect *Rectangle, factor int) {
	rect.width *= factor
	rect.height *= factor
}
func rectangleAreaB(rect Rectangle, factor int) {
	rect.width *= factor
	rect.height *= factor
}

func main() {
	var rect1 Rectangle
	var rect2 *Rectangle = new(Rectangle)
	rect3 := NewRectangle(20, 10)
	rect4 := &Rectangle{30, 20}
	rect5 := Rectangle{30, 30}

	rect1.height = 20
	rect2.height = 62

	fmt.Println(rect1)
	fmt.Println(rect2)
	fmt.Println(rect3)
	fmt.Println(rect4)

	area := rectangleArea(&rect5)
	fmt.Println(area)

	rect6 := Rectangle{30, 30}
	rectangleAreaA(&rect6, 10)
	fmt.Println(rect6)

	rect7 := Rectangle{30, 30}
	rectangleAreaB(rect7, 10)
	fmt.Println(rect7)

}
