package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	var a bool = true // Boolean
	fmt.Println("Boolean: ", a)

	var b int = 5 // Integer
	fmt.Println("Integer: ", b)
	fmt.Println(math.MinInt)
	fmt.Println(math.MaxInt)
	fmt.Println(bits.OnesCount8(math.MaxUint8))
	fmt.Println(bits.OnesCount8(math.MaxInt8))

	var myUint uint = 10
	fmt.Println(myUint)

	// bytes is alias of uint8
	// byte is mapping in ascii
	var myByte byte = 'A'
	fmt.Println(myByte)

	var c float32 = 3.14 // Floating point number
	fmt.Println("Float:   ", c)

	var myComplex = 10 + 1i
	fmt.Println(myComplex)

	var d string = "Hi!" // String
	fmt.Println("String:  ", d)

	var myString string = "Nhẫn"
	for i := 0; i < len(myString); i++ {
		fmt.Printf("%c", myString[i])
	}
	fmt.Println()

	myRune := []rune(myString)
	for i := 0; i < len(myRune); i++ {
		fmt.Printf("%c", myRune[i])
	}
	fmt.Println()

	var myRune2 rune = 'A'
	fmt.Printf("%T", myRune2)
}
