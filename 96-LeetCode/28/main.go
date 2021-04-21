package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("aaa", "aaa"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	var ri int = -1
	if haystack != "" && needle != "" {
		for i := 0; i < len(haystack); i++ {
			for j := 0; j < len(needle); {
				if haystack[i+j] == needle[j] {
					j++
					if j == len(needle) {
						ri = i
						return ri
					}
					if i+j == len(haystack) {
						return -1
					}
				} else {
					break
				}
			}
		}
	}
	return ri
}
