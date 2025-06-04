package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("virtual cpus =", runtime.NumCPU())
	var array []int
	fmt.Println(array)        // nil array has empty array as zero value
	fmt.Println(array == nil) // array is nil
	array = append(array, 1)
	fmt.Println(array)
	fmt.Println([]int{} == nil) // explicit empty array is not nil
	var object interface{}
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	json.Unmarshal(b, &object)
	fmt.Println(object)

	myChan := make(chan int)
	go func() {
		for i := range 5 {
			myChan <- i
		}
		close(myChan)
	}()
	for i := range myChan {
		fmt.Println("read channel using range", i)
	}
	ch := make(chan int)
	go func() {
		fmt.Println("ch read", <-ch)
	}()
	ch <- 0

	myChan <- 2 // panic: send on closed channel
}
