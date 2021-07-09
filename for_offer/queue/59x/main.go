package main

import (
	"fmt"
	"reflect"
)

func main() {
	for _, c := range "cdbed" {
		fmt.Printf("%s	%t %T\n", reflect.TypeOf(c), c, c)
	}
	// fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	res := []int{}
	if n == 0 || k == 0 {
		return res
	}
	left, right := 1-k, 0
	// 单调队列
	que := []int{}
	for right < n {
		// 移除队首
		if left > 0 && que[0] == nums[left-1] {
			que = que[1:]
		}
		for len(que) != 0 && que[len(que)-1] < nums[right] {
			que = que[:len(que)-1]
		}
		que = append(que, nums[right])
		if left >= 0 {
			res = append(res, que[0])
		}
		left++
		right++
	}
	return res
}
