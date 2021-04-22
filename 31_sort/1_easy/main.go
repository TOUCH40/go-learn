package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 9, 5, 7, 3, 5, 2, 9, 4, 6, 10}
	// fmt.Println(selectionSort(arr))
	fmt.Println(QuickSort(arr))
}

//归并排序
func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	i := length / 2
	left := MergeSort(arr[0:i])
	right := MergeSort(arr[i:])
	res := merge(left, right)
	return res
}

//合并数组
func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0
	l, r := len(left), len(right)
	//比较两个数组，谁小把元素值添加到结果集内
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
		} else {
			result = append(result, left[m])
			m++
		}
	}
	//如果有一个数组比完了，另一个数组还有元素的情况，则将剩余元素添加到结果集内
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result
}

//快排，以第一个值为标准，小于此值的放左边，大于此值放右边，将第一个值放中间，在分好的数组里如此往复
func QuickSort(arr []int) []int {

	return arr
}
