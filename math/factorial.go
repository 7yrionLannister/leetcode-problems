package main

import (
	"fmt"
	"math/big"
)

func main() {
	// for i := range 31 {
	// 	fmt.Println(i, ":", trailingZeroes(i))
	// }
	fmt.Println(30, ":", trailingZeroes(30))
}

// func trailingZeroes(n int) int {
// 	factorial := 1
// 	for i := range n {
// 		factorial *= (i + 1)
// 	}
// 	return countTrailingZeroes(factorial)
// }

// func countTrailingZeroes(n int) int {
// 	count := 0
// 	nStr := strconv.Itoa(n)
// 	for i := len(nStr) - 1; i > 0; i-- {
// 		if nStr[i] != '0' {
// 			break
// 		}
// 		count++
// 	}
// 	return count
// }

func trailingZeroes(n int) int {
	factorial := big.NewInt(1)
	var n64 int64 = int64(n)
	for i := int64(0); i < n64; i++ {
		factorial.Mul(factorial, big.NewInt(i+1))
	}
	fmt.Println("Factorial of", n, "is", factorial)
	return countTrailingZeroes(factorial)
}

func countTrailingZeroes(n *big.Int) int {
	count := 0
	nStr := n.String()
	for i := len(nStr) - 1; i > 0; i-- {
		if nStr[i] != '0' {
			break
		}
		count++
	}
	return count
}

// func trailingZeroes(n int) int {
// 	count := 0
// 	for n > 0 {
// 		n /= 5
// 		count += n
// 	}
// 	return count
// }
