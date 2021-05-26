package main

import (
	"fmt"
)

func main() {
	// in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(printNumbers(3))
}

func printNumbers(n int) []int {
	out := []int{}
	max := 0
	for i := 0; i < n; i++ {
		max = max*10 + 9
	}

	for i := 1; i <= max; i++ {
		out = append(out, i)
	}
	return out
}

// func printNumbers(n int) []int {
// 	out := []int{}
// 	if n == 0 {
// 		return out
// 	}
// 	tem := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	out = tem
// 	for i := 1; i < n; i++ {
// 		tem2 := []int{}
// 		for _, v := range tem {
// 			for j := 0; j < 10; j++ {
// 				tem2 = append(tem2, v*10+j)
// 			}
// 		}
// 		tem = tem2
// 		out = append(out, tem...)
// 	}
// 	return out
// }
