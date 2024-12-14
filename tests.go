package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	myCopyOfSlice := make([]int, len(mySlice))
	copy(myCopyOfSlice, mySlice)
	myCopyOfSlice[0] = 10
	println(mySlice)
	println(myCopyOfSlice)
	fmt.Println(mySlice[0])
	fmt.Println(myCopyOfSlice[0])
	fmt.Println("From index 1 to 3:", mySlice[1:4])
	fmt.Println("From index 1:", mySlice[1:])
	fmt.Println("Until index 3:", mySlice[:4])
	fmt.Println(getCoordinates())
	// Below, because Wheel is an anonymous struct, you have to redeclare it to access its fields.
	// With named structs like Chassis, you can access its fields directly.
	fmt.Println(Car{Name: "Ford XXX", Color: "Red", Brand: "Ford", Wheel: struct {
		Radius   int
		Material string
	}{Radius: 10, Material: "Aluminum"}, Chassis: Chassis{Material: "Steel"}})
	r := Rectangle{Width: 10, Height: 20}
	var s Shape = r // r is a Rectangle, Rectangle implements the Shape interface, so r is a Shape
	fmt.Println("area of r:", r.area())
	fmt.Println(s)
	//fmt.Println(area(r)) you can't call area() independently of Rectangle, because it is a method of Rectangle
}

// Named return values tell you what to expect from the function.
// If you don't name them, It is not clear what you are expecting (y, x) or (x, y)
// In this case, it is clear that you are expecting x and y to be returned.
func getCoordinates() (x, y int) {
	x = 10
	y = 20
	return // same as return x, y
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

type Shape interface {
	area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) area() float64 { // this function is a method of Rectangle
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return c.Radius * c.Radius * 3.14159
}

type MyCustomError struct {
	Message string
}

// instead of implementing the Error interface, we can use errors.New() to create a new error
// that way we avoid the need to implement the Error interface with a custom type
func (m MyCustomError) Error() string { // here we implement the Error interface
	return "MyCustomError: " + m.Message
}
