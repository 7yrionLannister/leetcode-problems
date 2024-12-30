package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	ticker := time.Tick(2 * time.Second)
	timeout := time.After(10 * time.Second)

	go func() {
		for {
			c1 <- "Every 500ms with for loop"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	stop := false
	for !stop {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case t := <-ticker:
			fmt.Println("Every 2 seconds with time.Tick", t)
		case end := <-timeout:
			fmt.Println("10 seconds have passed", end)
			stop = true
		}
	}
}
