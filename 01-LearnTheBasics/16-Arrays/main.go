package main

import "fmt"

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	array := [...]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator
	fmt.Println(x)
}
