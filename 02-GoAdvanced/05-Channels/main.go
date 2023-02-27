package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}

	messages := make(chan string)

	go func() {
		// write to channel messages
		messages <- "ping"
	}()

	// read from channel messages
	msg := <-messages
	fmt.Println(msg)

	done := make(chan bool)
	go func(done chan bool) {
		fmt.Println("Hello world goroutines")
		done <- true
	}(done)
	<-done
	fmt.Println("Done hello world")

	sendc := make(chan int)
	go func(sendc chan<- int) {
		sendc <- 10
	}(sendc)
	fmt.Println(<-sendc)

	ch := make(chan int)
	go func(chnl chan int) {
		for i := 0; i < 10; i++ {
			chnl <- i
		}
		close(chnl)
	}(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received: ", v, ok)
	}

	ch2 := make(chan int)
	go func(chnl chan int) {
		for i := 0; i < 10; i++ {
			chnl <- i
		}
		close(chnl)
	}(ch2)
	for v := range ch2 {
		fmt.Println("Received: ", v)
	}

	ch3 := make(chan string, 2)
	ch3 <- "buffered"
	ch3 <- "channel"
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
}
