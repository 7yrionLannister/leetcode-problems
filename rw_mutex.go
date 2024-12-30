package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0
	var m sync.RWMutex

	// goroutine 1
	// Since we are only reading data here, we can call the `RLock`
	// method, which obtains a read-only lock
	go func() {
		m.RLock()
		defer m.RUnlock()
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	// goroutine 2
	go func() {
		m.RLock()
		defer m.RUnlock()
		nIsPositive := n > 0
		time.Sleep(5 * time.Millisecond)
		if nIsPositive {
			fmt.Println(n, " is positive")
			return
		}
		fmt.Println(n, "is not positive")
	}()

	// goroutine 3
	// Since we are writing into data here, we use the
	// `Lock` method, like before
	go func() {
		m.Lock()
		n++
		m.Unlock()
	}()

	time.Sleep(time.Second)
}
