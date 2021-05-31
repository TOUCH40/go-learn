package main

import "fmt"

func main() {
	ns := []int{5, 7, 7, 8, 8, 10}

	fmt.Println(search(ns, 8))
	fmt.Println(search(ns, 6))
}
func search(nums []int, target int) int {
	return (findRight(nums, target) - findRight(nums, target-1))
}

// 计算有边界
func findRight(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
