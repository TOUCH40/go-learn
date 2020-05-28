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

}
