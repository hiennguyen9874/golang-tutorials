package main

import "fmt"

type Student struct {
	id   int
	name string
}

// Define method
// func (t Type) methodName(params) returns { //code }
// (t Type) => receiver
// 1. Value receiver
// 2. Pointer receiver
func (s Student) getName() string {
	return s.name
}

func (s Student) changeName() {
	fmt.Printf("*p changeName = %p\n", &s)
	s.name = "A"
}

func (s *Student) changeName2() {
	fmt.Printf("*p changeName2 = %p\n", s)
	s.name = "A"
}

func main() {
	st1 := Student{1, "Hien"}
	fmt.Println(st1)
	fmt.Printf("*p outside = %p\n", &st1)

	st1.changeName()
	fmt.Println(st1)

	st1.changeName2()
	fmt.Println(st1)
}
