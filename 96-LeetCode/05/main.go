package main

import "fmt"

func main() {
	In := []int{-1, 0, 1, 2, -1, -4}
	Out := threeSum(In)
	fmt.Println(Out)
}

// make(slice,cabpabity)
// sort+double pointer
func threeSum(nums []int) [][]int {
    out:=make([][]int,0)
    outMap:=make(map[int]int)
    for _,num:=range nums{
        outMap[num]=outMap[num]+1
    }
	
	return out
}
