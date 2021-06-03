package main

import "sort"

func main() {
	isStraight([]int{11, 8, 12, 8, 10})
}

func isStraight(nums []int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	k, gap := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			k++
			if k > 2 {
				return false
			}
		} else {
			if nums[i+1] == nums[i] {
				return false
			}
			gap += nums[i+1] - nums[i] - 1
		}
	}
	return k == gap
}
