package main

import "fmt"

func main() {
	// Short Declaration
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	// Multiple variable declarations with inferred types
	var firstName, lastName, age, salary = "John", "Maxwell", 28, 50000.0

	fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T\n",
		firstName, lastName, age, salary)
}
