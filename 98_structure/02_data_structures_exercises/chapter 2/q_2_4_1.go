package main

import "fmt"

func main() {
	orignArray := []int{1, 5, 2, 13, 5}
	pos := findMin(orignArray)
	fmt.Println(pos)
}

func findMin(slice []int) int {
	ret := 0
	if len(slice) > 0 {
		for i := 1; i < len(slice); i++ {
			if slice[i] > slice[i-1] {
				ret = i
			}
		}
	}

	return ret + 1
}
