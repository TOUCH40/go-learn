package main

import (
	"fmt"

	"sourcegraph.com/sqs/goreturns/returns"
)

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	if(len(matrix)==0){
		return []int{}
	}
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	res:=[]int{}
	for {
		for 
	}
	return res
}
