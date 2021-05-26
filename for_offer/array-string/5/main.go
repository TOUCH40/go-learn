package main

import (
	"fmt"
)

func main() {
	in := []int{1, 2, 3, 4}
	fmt.Println(exchange(in))
}

func exchange(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		for nums[l]%2 == 1 && l < r {
			l++
		}
		for nums[r]%2 == 0 && l < r {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}
