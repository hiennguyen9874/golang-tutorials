package main

import "fmt"

func addItem(item int, list ...int) {
	// list: slice => [2, 3, 4, 5, 6]
	list = append(list, item)
	fmt.Println(list)
}

func change(list ...int) {
	list[0] = 999
}

func main() {
	addItem(1, 2, 3, 4, 5, 6)

	var list = []int{100, 200, 300, 400, 500}
	addItem(500, list...)

	// Use list as slice to pass into function
	change(list...)
	fmt.Println(list)
}
