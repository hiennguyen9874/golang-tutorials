package main

import "os"

func main() {
	panic("a problem 1")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	panic("a problem 2")
}
