package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates(S string) string {
	var a = ""
	for _, _s := range S {
		a += _s
	}
	return a
}
