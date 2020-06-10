package main

import (
	"fmt"
)

func main() {
	In := []string{"a==b", "b!=c", "c==a"}
	ret := equationsPossible(In)
	fmt.Println(ret)
}

func equationsPossible(equations []string) bool {
	p := [26]*[]int{}
	for i := range p {
		p[i] = &[]int{i}
	}
	for _, s := range equations {
		i, j := int(s[0]-97), int(s[3]-97)
		if s[1] == '=' && p[i] != p[j] {
			if len(*p[i]) < len(*p[j]) {
				i, j = j, i
			}
			*p[i] = append(*p[i], *p[j]...)
			for _, k := range *p[j] {
				p[k] = p[i]
			}
		}
	}

	for _, s := range equations {
		i, j := int(s[0]-97), int(s[3]-97)
		if (s[1] == '=') != (p[i] == p[j]) {
			return false
		}
	}
	return true
}
