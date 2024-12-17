package main

import (
	"fmt"
	"strings"
	"sync"
	//"maps"
	//"math"
	//"strconv"
	//theMapsPackage "maps" you can import a package with a different name
)

// iota is a special constant that starts at 0 and increments by 1 for each new constant
const (
	myConst1 = iota // myConst1 = 0
	myConst2        // myConst2 = 1
	myConst3        // myConst3 = 2
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	myCopyOfSlice := make([]int, len(mySlice))
	copy(myCopyOfSlice, mySlice)
	myCopyOfSlice[0] = 10
	fmt.Println(mySlice[0])
	fmt.Println(myCopyOfSlice[0])
	fmt.Println("From index 1 to 3:", mySlice[1:4])
	fmt.Println("From index 1:", mySlice[1:])
	fmt.Println("Until index 3:", mySlice[:4])
	demoSlice := make([]int, 5)
	demoSlice = append(demoSlice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Demo slice:", demoSlice)
	fmt.Println("Length of demoSlice:", len(demoSlice))   // number of elements
	fmt.Println("Capacity of demoSlice:", cap(demoSlice)) // number of elements that can be stored before a reallocation
	fmt.Println(getCoordinates())
	// Below, because Wheel is an anonymous struct, you have to redeclare it to access its fields.
	// With named structs like Chassis, you can access its fields directly.
	fmt.Println(Car{Name: "Ford XXX", Color: "Red", Brand: "Ford", Wheel: struct {
		Radius   int
		Material string
	}{Radius: 10, Material: "Aluminum"}, Chassis: Chassis{Material: "Steel"}})
	r := &Rectangle{Width: 10, Height: 20}
	r.area()
	// any == inferface{} empty interface
	// all types in go implement the empty interface, so you can use any type as a function argument
	// the empty interface (this type is any) is the type with no methods and no fields, it is similar to Object in Java
	var s Shape = r // r is a Rectangle, Rectangle implements the Shape interface, so r is a Shape
	fmt.Println("area of r:", r.area())
	fmt.Println(s)
	//fmt.Println(area(r)) you can't call area() independently of Rectangle, because it is a method of Rectangle

	myStrings := []string{"a", "b", "c"}
	fmt.Println("Printing the slice through variatic function")
	variaticFunction(myStrings...) // ... in the caller acts as a spread operator
	// this could also be written as:
	variaticFunction("c", "d", "e")
	myMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 0}
	// non-existent elements are initialized to the zero value of the map's key type
	// to differentiate the missing key from the actual zero value, use the ok result of lookup, which tells you whether the key existed
	existing0Mapping, ok := myMap["d"]
	fmt.Println("Existing element of myMap:", existing0Mapping, ", ok:", ok)
	nonExistingElement, nok := myMap["e"]
	fmt.Println("Non-existing element of myMap:", nonExistingElement, ", nok:", nok)
	fmt.Println("Keys of myMap:", myMap)

	myFunction := FunctionProducer()
	fmt.Println("Calling myFunction:", myFunction(10))

	// pointers
	var x int = 10
	var y *int = &x // y is a pointer to x, which means that y points to the same memory location as x
	// y is not a copy of x, it is a reference to x
	// if you change the value of x, y will also change. The oposite is also true, if you change the value of y, x will change
	*y = 20 // dereferencing y (*y), which means that we are accessing the value that y points to
	println("x =", x)
	println("y = &x =", *y)
	x = 15
	strings.ReplaceAll("Hello World", "World", "Universe") // strings package
	println("Memory location of x:", &x, ", value of x:", x)
	println("Memory location of y:", &y, ", value of y:", *y, ", address to which y points:", y)

	// CHANNELS
	myChannel := make(chan int) // create a channel of type int
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Sending", i, "to the channel")
			myChannel <- i // send i to the channel
		}
		close(myChannel) // close the channel to signal the consumer that no more values will be sent
		// IMPORTANT: always close the channel when you are done sending values (usually at the end of the goroutine)
	}()

	// for i := range myChannel {
	// 	fmt.Println("Receiving from the channel using range:", i)
	// }
	readChannel(myChannel)
	callGenericFunction()
	var ordered1 int = 10
	var ordered2 int = 20
	fmt.Println("ordered1 < ordered2:", isLessThan(ordered1, ordered2))

	//CLOSURES
	// We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.
	nextInt := intSeq()

	// See the effect of the closure by calling nextInt a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())

	// INTERFACES
	printArea(r)
	c := &Circle{Radius: 10}
	printArea(c)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func readChannel(myChannel <-chan int) { // you can make a channel read-only by using the <-chan syntax
	for i, ok := <-myChannel; ok; i, ok = <-myChannel { // receive from the channel, and check if it is still open and not empty
		fmt.Println("Receiving from the channel using ok indicator:", i)
	}
}

// MUTEXES are used to protect shared resources from concurrent access, like a database connection or a file.
var mux sync.Mutex = sync.Mutex{} // create a mutex
func protected() {
	mux.Lock()
	defer mux.Unlock()
	fmt.Println("Protected section")
}

// Named return values tell you what to expect from the function.
// If you don't name them, It is not clear what you are expecting (y, x) or (x, y)
// In this case, it is clear that you are expecting x and y to be returned.
func getCoordinates() (x, y int) {
	x = 10
	y = 20
	return // same as return x, y
}

// just as a function can receive other functions as arguments, it can return other functions
func FunctionProducer() func(int) int {
	return func(x int) int {
		return x * x
	}
}

type Car struct {
	Name  string
	Color string
	Brand string
	Wheel struct { // anonymous struct, only used inside Car. In general, prefer named structs.
		Radius   int
		Material string
	}
	Chassis Chassis
}

type Chassis struct { // named struct. In general, prefer named structs.
	Material string
}

// INTERFACES
type Shape interface {
	area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) area() float64 { // this function is a method of Rectangle
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c *Circle) area() float64 { // prefer using pointer receivers to avoid copying the struct
	return c.Radius * c.Radius * 3.14159
}

// this is the same as above, but with a value receiver
// func (c Circle) area() float64 {
// 	return c.Radius * c.Radius * 3.14159
// }

func printArea(s Shape) {
	fmt.Println("area of s:", s.area())
}

type MyCustomError struct {
	Message string
}

// instead of implementing the Error interface, we can use errors.New() to create a new error
// that way we avoid the need to implement the Error interface with a custom type
func (m MyCustomError) Error() string { // here we implement the Error interface
	return "MyCustomError: " + m.Message
}

func variaticFunction(a ...string) { // ... in the argument list acts as an indicator of a variadic argument (arbitrary number of arguments of the same type)
	fmt.Println("Printing all the arguments", a)
}

func callGenericFunction() {
	MyGeneric[int]{Value: 10}.Print()
	// generic type parameter can be inferred from the function call
	fmt.Println("Generic function:", genericFunction(15))              // generic function call without type parameter
	fmt.Println("Generic function:", genericFunction[string]("Hello")) // generic function call with type parameter
}

// generics
type MyGeneric[T any] struct {
	Value T
}

func (m MyGeneric[X]) Print() {
	fmt.Println("Printing the value of the generic:", m.Value)
}

func genericFunction[T any](x T) T { // instead of any, you can specify another interface. It would work similar to subtyping in Java
	return x
}

// with generics, a new way of defining interfaces is possible
// all this types implement the Ordered interface, so we can use it as a type constraint
// these types support the comparison operators: <, <=, >, >=, ==, !=
type Ordered interface {
	~int | ~uint | ~float64 | ~float32 | ~uint16 | ~uint32 | ~uint64 | ~int16 | ~int32 | ~int64 | ~uintptr | ~string
}

func isLessThan[T Ordered](a, b T) bool {
	return a < b
}

func vvv[S ~[]C, C comparable](com1, com2 C) bool {
	return com1 == com2
}
