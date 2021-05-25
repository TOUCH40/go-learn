package main

import (
	"fmt"
)

func main() {
	str := "We are happy."
	res := ""
	for _, r := range str {
		if r == ' ' {
			res += "%20"
		} else {
			res += string(r)
		}
	}
	// str = strings.ReplaceAll(str, " ", "%20")
	fmt.Println(res)
}
