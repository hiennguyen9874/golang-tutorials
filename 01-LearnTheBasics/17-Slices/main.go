package main

import "fmt"

func main() {
	var slice []int
	fmt.Println(slice)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	a := []int{1, 2, 3}
	b := a[:2]
	b[0] = 9
	fmt.Println(a)
	fmt.Println(b)

	countries := [...]string{"VN", "USA", "CANADA", "FRANCE", "CHINA"}
	slice10 := countries[2:3]
	fmt.Println(slice10)

	// Len: so luong phan tu cua slice
	fmt.Println(len(slice10))

	// Cap: So luong phan tu cua underlying array bat dau tu vi tri start khi ma slice duoc tao
	fmt.Println(cap(slice10))
}
