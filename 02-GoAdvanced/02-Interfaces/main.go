package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

// Value receiver
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// Pointer receiver
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Movement interface {
	move()
}

type Animal interface {
	speak()
}

type Dog struct{}

func (d Dog) speak() {
	fmt.Println("Go go go")
}

func (d Dog) move() {
	fmt.Println("dog move")
}

type NextAnimal interface {
	Movement
	Animal
}

func printType(i interface{}) {
	fmt.Printf("%T\n", i)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())

	dog := Dog{}

	var m Movement = dog
	m.move()

	var an Animal = dog
	an.speak()

	var na NextAnimal = dog
	na.move()
	na.speak()

	printType(1)
	printType(1.1)
	printType("A")
	printType(struct {
		id   int
		name string
	}{1, "A"})
}
