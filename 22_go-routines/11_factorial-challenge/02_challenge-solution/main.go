package main

import (
	"fmt"
)

func main() {
	res := factorial(4)
	// for r := range res {
	// 	fmt.Println(r)
	// }
	fmt.Println(<-res)
}

func factorial(num int) chan int {
	out := make(chan int)
	go func() {
		n := num
		sum := 1
		for i := 1; i <= n; i++ {
			sum = i * sum
		}
		out <- sum
		// close(out)
	}()
	return out
}
