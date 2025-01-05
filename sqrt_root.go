package main

import "fmt"

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	i, j := 0, x
	nearestSqrt := 0
	for i != j {
		num := (i + j) / 2
		num2 := num * num
		if num2 == x {
			nearestSqrt = num
			i = j // break
		} else if num2 < x {
			if num2 > nearestSqrt*nearestSqrt {
				nearestSqrt = num
			} else {
				i = num + 1
			}
		} else {
			j = num - 1
		}
	}
	return nearestSqrt
}
