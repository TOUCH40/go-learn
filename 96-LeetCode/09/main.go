package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(102201))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	it := 0
	ix := x
	for ix > 0 {
		it = it*10 + ix%10
		ix = ix / 10
	}
	return it == x
}
