package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	txt := "is2 sentence4 This1 a3"
	arr := strings.Split(txt, " ")
	dic := make(map[int]string, len(arr))
	for _, item := range arr {
		index, _ := strconv.Atoi(item[len(item)-1:]) //int(idd)
		dic[index] = item[0 : len(item)-1]
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = dic[i-1]
	}
	c := strings.Join(arr, " ")
	fmt.Println(c)
}
