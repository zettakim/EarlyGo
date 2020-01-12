package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

func formatString(arg interface{}) string {
	switch arg.(type) {
	case int:
		i := arg.(int)
		return strconv.Itoa(i)
	case float32:
		f := arg.(float32)
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
	case float64:
		f := arg.(float64)
		return strconv.FormatFloat(f, 'f', -1, 64)
	case string:
		s := arg.(string)
		return s
	case Person:
		p := arg.(Person)
		return p.name + " " + strconv.Itoa(p.age)
	case *Person:
		p := arg.(*Person)
		return p.name + " " + strconv.Itoa(p.age)
	default:
		return "Error"
	}
}
func main() {
	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println("Hello, world!")
	fmt.Println()

	fmt.Println(formatString(Person{"Maria", 20}))
	fmt.Println(formatString(&Person{"John", 12}))

	var andrew = new(Person)
	andrew.name = "Andrew"
	andrew.age = 35

	fmt.Println(formatString(andrew))

	var t interface{}
	t = Person{"Zetta", 53}

	if v, ok := t.(Person); ok {
		fmt.Println(v, ok)
	}
}
