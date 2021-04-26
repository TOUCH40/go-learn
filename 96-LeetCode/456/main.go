package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4}
	nb := find132pattern(num)
	fmt.Printf("%b", nb)
}

func find132pattern(nums []int) bool {
	N := len(nums)
	numsi := nums[0]
	for j := 1; j < N; j++ {
		for k := N - 1; k > j; k-- {
			if numsi < nums[k] && nums[k] < nums[j] {
				return true
			}
		}
		if numsi > nums[j] {
			numsi = nums[j]
		}
	}
	return false
}
