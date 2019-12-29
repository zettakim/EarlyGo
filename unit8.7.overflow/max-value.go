package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 uint8 = math.MaxUint8
	var num2 uint16 = math.MaxUint16
	var num3 uint32 = math.MaxUint32
	var num4 uint64 = math.MaxUint64
	// var num5 uint8 = math.MaxUint8 + 1 // 오버플로우
	// var num1 uint8 = 0 - 1  // 오버플로우 컴파일 에러 발생

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
}
