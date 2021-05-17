package main

import (
	"fmt"
)

func main() {

	c := memLeak(8, 11)
	fmt.Println(c)
}

func memLeak(memory1 int, memory2 int) []int {
	ret := []int{1, memory1, memory2}
	var seekfun func()
	seekfun = func() {
		maxindex := 1
		if ret[1] < ret[2] {
			maxindex = 2
		}
		if ret[maxindex] < ret[0] {
			return
		}
		ret[maxindex] -= ret[0]
		ret[0]++
		seekfun()
	}
	seekfun()

	return ret
}
