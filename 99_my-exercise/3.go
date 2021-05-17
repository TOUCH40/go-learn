package main

import (
	"fmt"
)

type FindSumPairs struct {
	nums1 map[int]int
	nums2 []int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	c := new(FindSumPairs)
	ma := make(map[int]int)
	for _, i := range nums1 {
		if ma[i] == 0 {
			ma[i] = 1
		} else {
			ma[i] += 1
		}
	}
	c.nums1 = ma
	c.nums2 = nums2
	return *c
}

func (this *FindSumPairs) Add(index int, val int) {
	this.nums2[index] += val
}

func (this *FindSumPairs) Count(tot int) int {
	sum := 0
	ma := make(map[int]int)
	for _, i := range this.nums2 {
		if ma[i] == 0 {
			ma[i] = 1
		} else {
			ma[i] += 1
		}
	}
	for i, j := range ma {
		if this.nums1[tot-i] != 0 {
			sum += (this.nums1[tot-i] + j)
		}
	}
	return sum
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
func main() {
	n1 := []int{1, 1, 1, 1, 1}
	n2 := []int{1, 1, 1, 1, 1}
	c := Constructor(n1, n2)
	c.Add(2, 2)
	c.Add(2, 2)
	c.Add(2, 2)
	fmt.Println(c.nums2)
}

type FindSumPairs struct{}

var (
	x, y []int
	cnt  map[int]int
)

func Constructor(nums1, nums2 []int) (_ FindSumPairs) {
	cnt = map[int]int{}
	for _, v := range nums2 {
		cnt[v]++
	}
	x, y = nums1, nums2
	return
}

func (FindSumPairs) Add(i, val int) {
	cur := y[i]
	cnt[cur]--
	cnt[cur+val]++
	y[i] += val
}

func (FindSumPairs) Count(tot int) (ans int) {
	for _, v := range x {
		ans += cnt[tot-v]
	}
	return
}
