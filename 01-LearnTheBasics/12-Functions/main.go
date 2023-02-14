package main

import "fmt"

// Function accepting arguments
func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// The return values of a function can be named in Golang
func rectangle(l int, b int) (area int) {
	var parameter int
	parameter = 2 * (l + b)
	fmt.Println("Parameter:", parameter)

	area = l * b
	return // Return statement without specify variable name
}

// Golang Functions Returning Multiple Values
func rectangle2(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

// Golang Passing Address to a Function
func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	return
}

// Anonymous Functions in Golang
var (
	area = func(l int, b int) int {
		return l * b
	}
)

// Returning Functions from other Functions
func sum(x, y int) int {
	return x + y
}
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func main() {
	// Simple function with parameters in Golang
	add(1, 2)

	// Simple function with return value in Golang
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	// The return values of a function can be named in Golang
	fmt.Println("Area:", rectangle(20, 30))

	// Golang Functions Returning Multiple Values
	var a, p int
	a, p = rectangle2(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)

	// Golang Passing Address to a Function
	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)

	update(&age, &text)

	fmt.Println("After :", text, age)

	// Anonymous Functions in Golang
	fmt.Println(area(20, 30))
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)

	// Closures Functions in Golang
	l := 20
	b := 30

	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()

	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}

	// Returning Functions from other Functions
	partial := partialSum(3)
	fmt.Println(partial(7))

	fmt.Println(squareSum(5)(6)(7))
}
