package main

import "fmt"

func main() {
	abc := []string{"aa", "a"}
	fmt.Println(longestCommonPrefix(abc))
}
func longestCommonPrefix(strs []string) string {
	out := ""
	strLen := len(strs)
	key := 0
	for {
		letter := ""
		if strLen <= 0 || len(strs[0]) <= key {
			letter = ""
			break
		} else {
			letter = string(strs[0][key])
		}
		for i := 0; i < strLen; i++ {
			if len(strs[i]) <= key || string(strs[i][key]) != letter {
				letter = ""
				break
			}
		}
		if letter == "" {
			break
		} else {
			out += letter
			key++
		}
	}
	return out
}
