package main

import "fmt"

func main() {
	In := []int{100, 4, 200, 1, 3, 2}
	Out := longestConsecutive(In)
	fmt.Println(Out)
}

// hash table
// map 判断 value,ok=map_num[key]
// map 赋值 map_num[key]=true
// for _(index),num(value)

func longestConsecutive(nums []int) int {
	longest := 0
	mapNum := make(map[int]bool)
	for _, num := range nums {
		mapNum[num] = true
	}
	for _, num := range nums {
		_, ok := mapNum[num-1]
		if !ok {
			thisLen := 1
			for {
				_, ok2 := mapNum[num+thisLen]
				if !ok2 {
					break
				}
				thisLen++
			}
			if thisLen > longest {
				longest = thisLen
			}
		}
	}
	return longest
}
