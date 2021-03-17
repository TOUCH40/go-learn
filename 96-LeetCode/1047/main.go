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
		if len(stack) != 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1] // 出栈
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}
