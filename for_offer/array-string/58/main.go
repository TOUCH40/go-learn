package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords2("a good   example"))
}

func reverseWords(s string) string {
	out := ""
	top := -1
	stack := []string{}
	for i := 0; i < len(s); {
		wordT := ""
		for i < len(s) && s[i] != ' ' {
			wordT += string(s[i])
			i++
		}
		if wordT != "" {
			stack = append(stack, wordT)
			top++
		}

		for i < len(s) && s[i] == ' ' {
			i++
		}
	}
	for top > -1 {
		out += stack[top]
		if top != 0 {
			out += " "
		}
		top--
	}
	return out
}
func reverseWords2(s string) string {
	s = strings.Trim(s, " ")
	out := ""
	l, r := len(s)-1, len(s)-1
	for l > -1 {
		for l > -1 && s[l] != ' ' {
			l--
		}
		out += s[l+1:r+1] + " "
		for l >= 0 && s[l] == ' ' {
			l--
		}
		r = l
	}
	return strings.Trim(out, " ")
}

1,2,3,4,5,6   2