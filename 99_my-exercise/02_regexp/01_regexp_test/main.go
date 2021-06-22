package main

import (
	"fmt"
	"regexp"
)

func main() {
	fileName := "01.课程介绍.vep"
	reg := fmt.Sprintf(`\.%s$`, "vep") //`\.vep$`
	ret := regexp.MustCompile(reg)
	if ret == nil {
		fmt.Println("not regexp")
		return
	}
	isVep, err := regexp.MatchString(reg, fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(isVep)
	fmt.Println(bubbleSort([]int{1, 3, 1, 5, 3, 7, 8, 3, 4, 65, 234, 22}))
}

func bubbleSort(arr []int) (out []int) {
	out = arr
	for i := 0; i < len(out); i++ {
		flag := false
		for j := len(out) - 1; j > i; j-- {
			if out[j] < out[j-1] {
				flag = true
				out[j], out[j-1] = out[j-1], out[j]
			}
		}
		if !flag {
			break
		}
	}
	return
}



