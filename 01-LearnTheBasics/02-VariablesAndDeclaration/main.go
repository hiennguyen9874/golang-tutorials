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
}
