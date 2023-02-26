package main

import (
	"fmt"
	"math"
)

var c, python, java bool

const s string = "constant"

func main() {
	var i int
	fmt.Println(i, c, python, java)

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	fmt.Println(s)

	const n = 500000000

	const g = 3e20 / n
	fmt.Println(g)

	fmt.Println(int64(g))

	fmt.Println(math.Sin(n))

	var (
		name    string
		address string
		age     int
	)
	name = "Hien"
	address = "Viet nam"
	age = 1

	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(age)

	var name2, address2, age2 = "Hien", "Vietnam", 1

	fmt.Println(name2)
	fmt.Println(address2)
	fmt.Println(age2)
}
