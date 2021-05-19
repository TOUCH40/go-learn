package main

import (
	"fmt"
)

func main() {
	a := []int{12, 3, 111, 23, 65, 45}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
func quickSort(a []int, left int, right int) {
	if left >= right { //一定是left >= right
		return
	}
	temp := a[left]
	start := left
	stop := right
	for right != left {
		for right > left && a[right] >= temp {
			right--
		}
		for left < right && a[left] <= temp {
			left++
		}
		if right > left {
			a[right], a[left] = a[left], a[right]
		}
	}
	a[right], a[start] = temp, a[right]
	quickSort(a, start, left)
	quickSort(a, right+1, stop)
}
