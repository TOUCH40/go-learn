package main

import (
	"fmt"
)

func main() {

	c := []int{3, 4, 5, 6, 7, 8}
	fmt.Println(subsetXORSum(c))
}

func subsetXORSum(nums []int) int {
	ret := 0
	subColl := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		for _, item := range subColl {
			subColl = append(subColl, item^nums[i])
		}
		subColl = append(subColl, nums[i])
	}

	for _, it := range subColl {
		ret += it
	}
	return ret
}
