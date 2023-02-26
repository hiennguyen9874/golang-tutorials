// Go program to find the
// average of numbers
package main

import "fmt"

func main() {

	// taking the required
	// data into variables
	var totalsum int = 846
	var number int = 19
	var avg float32

	// explicit type conversion
	avg = float32(totalsum) / float32(number)

	// Displaying the result
	fmt.Printf("Average = %f\n", avg)

	a := 6 / 3   // both are int, a is int
	f := 6.3 / 3 // float and int, f is float

	fmt.Println(a, f) // 2 2.1

	var myInt, myFloat = 1, 2.2
	fmt.Println(myInt + int(myFloat))
}
