package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates(S string) string {
	stack := []byte{}
	for i := range S {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
	// var a = make([]string)
	// i := 0
	// for _, _s := range S {
	// 	if len(a) == 0 || a[i] != _s {
	// 		i++
	// 		a[i] = _s
	// 	} else {
	// 		i--
	// 	}
	// }
	// for j:=1;j<i;j++
	// return a
}
