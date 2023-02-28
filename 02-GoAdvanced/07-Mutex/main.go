package main

import (
	"fmt"
	"sync"
	"time"
)

// a simple function that returns true if a number is even
func isEven(n int) bool {
	return n%2 == 0
}

type intLock struct {
	val int
	sync.Mutex
}

func (n *intLock) isEven() bool {
	return n.val%2 == 0
}

func main() {
	n := 0
	var m sync.Mutex

	// goroutine 1
	// reads the value of n and prints true if its even
	// and false otherwise
	go func() {
		m.Lock()
		defer m.Unlock()
		nIsEven := isEven(n)
		// we can simulate some long running step by sleeping
		// in practice, this can be some file IO operation
		// or a TCP network call
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, "is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	// goroutine 2
	// modifies the value of n
	go func() {
		m.Lock()
		defer m.Unlock()
		time.Sleep(time.Millisecond)
		n++
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)

	l := 0
	var k sync.RWMutex

	// goroutine 1
	// Since we are only reading data here, we can call the `RLock`
	// method, which obtains a read-only lock
	go func() {
		k.RLock()
		defer k.RUnlock()
		nIsEven := isEven(l)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(l, "is even")
			return
		}
		fmt.Println(l, "is odd")
	}()

	// goroutine 2
	go func() {
		k.RLock()
		defer k.RUnlock()
		nIsPositive := l > 0
		time.Sleep(5 * time.Millisecond)
		if nIsPositive {
			fmt.Println(l, "is positive")
			return
		}
		fmt.Println(l, "is not positive")
	}()

	// goroutine 3
	// Since we are writing into data here, we use the
	// `Lock` method, like before
	go func() {
		k.Lock()
		l++
		k.Unlock()
	}()

	time.Sleep(time.Second)

	o := &intLock{val: 0}

	go func() {
		o.Lock()
		defer o.Unlock()
		nIsEven := o.isEven()
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(o.val, "is even")
			return
		}
		fmt.Println(o.val, "is odd")
	}()

	go func() {
		o.Lock()
		o.val++
		o.Unlock()
	}()

	time.Sleep(time.Second)
}
