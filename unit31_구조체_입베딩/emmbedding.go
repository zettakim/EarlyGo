package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello~")
}

type Student struct {
	p      Person // has-a
	school string
	grade  int
}

type Student2 struct {
	Person // is-a
	school string
	grade  int
}

func main() {
	var s Student
	s.p.greeting()

	var s2 Student2

	s2.Person.greeting()
	s2.greeting()
}
