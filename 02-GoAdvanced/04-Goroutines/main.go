package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func g1() {
	fmt.Println("g1")
}

func main() {
	go say("world")
	say("hello")

	go g1()

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	ng := runtime.NumGoroutine()
	fmt.Println(ng)

	time.Sleep(1 * time.Second)
	fmt.Println("done")

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		say("Hello world")
	}()
	wg.Wait()
}
