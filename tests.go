package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/7yrionLannister/leetcode-problems/a/b"
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
	// fmt.Println(area(r)) you can't call area() independently of Rectangle, because it is a method of Rectangle

	myStrings := []string{"a", "b", "c"}
	fmt.Println("Printing the slice through variatic function")
	variaticFunction(myStrings...) // ... in the caller acts as a spread operator
	// this could also be written as:
	variaticFunction("c", "d", "e")
	myMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 0}
	randonNum, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println("Random number:", randonNum)
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
	myChannel := make(chan int) // create a channel of type int (unbuffered, synchronous)
	go func() {
		defer close(myChannel) // close the channel to signal the consumer that no more values will be sent
		// IMPORTANT: for unbuffered channels always close the channel when you are done sending values (usually at the end of the goroutine)
		for i := 0; i < 10; i++ {
			fmt.Println("Sending", i, "to the unbuffered channel")
			myChannel <- i // send i to the channel
		}
	}()

	readChannel(myChannel)

	// enterito := <- myChannel
	// anotherChannel <- enterito
	// anotherChannel <- <- myChannel // this is the same as the two lines above

	myBufferedChannel := make(chan int, 3) // create a buffered channel of type int with a buffer of 3 elements (asynchronous)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("Sending", i, "to the buffered channel")
			myBufferedChannel <- i // send i to the channel
		}
		close(myBufferedChannel)
	}()

	for i := range myBufferedChannel {
		fmt.Println("Receiving from the buffered channel using range:", i)
	}

	// ch := make(chan int)
	// ch <- 2 // this blocks indefinitely because there are not receivers for ch

	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	// messages <- "deadlock" // this will cause a deadlock, because the channel is full
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// unbufferedChannel := make(chan int)
	// unbufferedChannel <- 1 // this blocks until there is a receiver for the channel. As there is no receiver, this will cause a deadlock
	// unbufferedOrBufferedChannel := make(chan int, 1) // this is a buffered channel with a buffer of 1 element
	// <-unbufferedOrBufferedChannel // this blocks until a message is sent to the channel, no matter if it is buffered or unbuffered. As the channel is empty, this will cause a deadlock
	// var once sync.Once
	// once.Do(func() {
	// 	once.Do(func() {
	// 		fmt.Println("Deadlock because the last call (this one) to [Do] locks until the first [Do] (the outer one) returns")
	// 	})
	// })

	c1 := make(chan string)
	c2 := make(chan string)
	var atomicInt32 atomic.Int32
	atomicInt32.Store(20)
	atomicInt32.Add(2)
	atomicInt32.Add(5)
	fmt.Println("Atomic Int32 value:", atomicInt32.Load())

	select {
	case msg1 := <-c1:
		fmt.Println("[not executed] received with select", msg1)
	case msg2 := <-c2:
		fmt.Println("[not executed] received with select", msg2)
	case t := <-time.After(1 * time.Second): // time.After returns a channel that will send a value after the specified time
		fmt.Println("timeout, current time is:", t)
	}

	// Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We’ll use select to await both of these values simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received with select", msg1)
		case c1 <- "coco": // you can also send a value to a channel with select
			fmt.Println("sent with select")
		case msg2 := <-c2:
			fmt.Println("received with select", msg2)
			// default: // the default case is executed if no other case is ready
			// 	fmt.Println("default case, no message received")
		}
	}

	// GENERICS
	callGenericFunction()
	var ordered1 int = 10
	var ordered2 int = 20
	fmt.Println("ordered1 < ordered2:", isLessThan(ordered1, ordered2))

	// CLOSURES
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

	var forma Forma = r
	forma.area()
	printAreaGeneric[*Circle](c)
	printAreaGeneric(r)

	// ERRORS

	// with convertion you can check if an error is of a certain type, with equality you can check if it is a specific error
	// if e, ok := err.(*MyCustomError); ok && e.Err == ErrPermission {
	// 	// query failed because of a permission problem
	// }

	// there is a better approach to handle errors in Go
	// this checks if the error is of a certain type, or if it wraps an error of that type (or any other error that wraps it)
	// it goes through the error chain until it finds the error equivalent to the one you are looking for
	// if errors.Is(err, ErrPermission) {
	// 	// err, or some error that it wraps, is a permission problem
	// }

	// var e *MyCustomError
	// Note: *MyCustomError is the type of the error.
	// errors.As examines the tree of its first argument looking for an error that can be assigned (casted) to its
	// second argument, which must be a pointer. If it succeeds, it performs the assignment and
	// returns true. Otherwise, it returns false
	// this is similar to if e, ok := err.(*MyCustomError); ok && e.Err == ErrPermission {}
	// if errors.As(err, &e) {
	// 	// err is a *MyCustomError, and e is set to the error's value
	// }

	// f, err := os.Open(filename)
	// if err != nil {
	// 	// The *os.PathError returned by os.Open is an internal detail.
	// 	// To avoid exposing it to the caller, repackage it as a new
	// 	// error with the same text. We use the %v formatting verb, since
	// 	// %w would permit the caller to unwrap the original *os.PathError.
	// 	return fmt.Errorf("%v", err)
	// }

	// MUTEXES
	var wg sync.WaitGroup
	myVal := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go protected(&wg, &myVal)
	}
	wg.Wait()
	// var iiiii atomic.Value
	fmt.Println(b.SayHello()) // b is in a package that is inside the a package, so the import path is "github.com/7yrionLannister/leetcode-problems/a/b"
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func readChannel(myChannel <-chan int) {
	// you can make a channel read-only by using the <-chan syntax
	// sending to a read-only channel will cause a compilation error
	// myChannel <- 10
	// you can make a channel write-only by using the chan<- syntax
	// receiving from a write-only channel will cause a compilation error
	// i := <-myChannel
	for i, ok := <-myChannel; i < 10 && ok; i, ok = <-myChannel { // receive from the channel, and check if it is still open and not empty
		fmt.Println("Receiving from the unbuffered channel using ok indicator:", i)
	}
}

// MUTEXES are used to protect shared resources from concurrent access, like a database connection or a file.
var mux sync.Mutex // create a mutex
func protected(wg *sync.WaitGroup, myVal *int) {
	defer wg.Done()
	mux.Lock()
	defer mux.Unlock()
	*myVal++ // this is the protected resource
	fmt.Println("Protected section", *myVal)
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

// you can also define an interface that extends another interface
type Forma interface {
	Shape
}

// you can also define an interface that is a union of other interfaces.
// the only constraint is that the interfaces must not have methods.
// you can't define a union of interfaces with methods, because it would be ambiguous.
// however, the interfaces in the union can have other interface extensions (like [OtherInterface]).
// this is useful when you want to define a function that can receive any of the interfaces in the union.
// it's useful as a type constraint, because you can't create a variable of an interface union type.
type InterfaceUnion interface {
	OneInterface | OtherInterface
}

type (
	OneInterface   interface{}
	OtherInterface interface {
		AnotherInterface | YetAnotherInterface
	}
)

type (
	AnotherInterface    interface{}
	YetAnotherInterface interface {
		TheLastInterface
	}
)
type TheLastInterface interface{}

// the tilde (~) is used to define a type constraint that is satisfied by underlying types (int but also anything that can be converted to int or embeds int)
type MyInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func printArea(s Shape) {
	fmt.Printf("s is a %T and its value is %+v\n", s, s) // %T prints the type of the variable
	switch v := s.(type) {                               // type switch
	// the type assertion below is a type switch, which is a switch statement that checks the type of the variable
	case *Rectangle:
		fmt.Printf("v is a %T and its value is %+v\n", v, v) // %T prints the type of the variable
		// the type assertion below is not necessary, because the type switch already checked the type and stored it in v as a Rectangle
		// the ok variable is true if the type assertion succeeded, false otherwise (if false, the variable is the zero value of the type)
		r, ok := s.(*Rectangle) // type assertion, extract the value from the interface into a variable of the concrete type
		if ok {
			fmt.Println("ok is true, type assertion to Rectangle succeeded")
		}
		fmt.Printf("s is a Rectangle: %+v\n", r)
	case *Circle:
		fmt.Printf("v is a %T and its value is %+v\n", v, v) // %T prints the type of the variable
		c := s.(*Circle)                                     // type assertion, extract the value from the interface into a variable of the concrete type
		fmt.Printf("s is a Circle: %+v\n", c)
	}
	fmt.Println("area of s:", s.area()) // type assertion
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

func printAreaGeneric[T Forma](s T) {
	fmt.Println("<generic> area of s:", s.area())
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
