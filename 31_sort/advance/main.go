package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 9, 5, 7, 3, 5, 2, 9, 4, 6, 10}
	// fmt.Println(selectionSort(arr))
	fmt.Println(BubbleSort(arr))
}

//选择排序
//思路：每次循环找出最小的数，跟数组第一个数交换顺序，接下来在剩余的数里重复以上逻辑
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		maxIndex := i
		for j := i; j < n; j++ {
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
		}
		// temp := arr[i]
		// arr[i] = arr[maxIndex]
		// arr[maxIndex] = temp
		arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
	}
	return arr
}

//插入排序，类似扑克牌起牌，将未排序的数据插入到已排序的数据中
func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] > arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	return arr
}

//插入排序优化版，用赋值代替交换操作
func InsertionSortPro(arr [11]int) [11]int {
	length := len(arr)
	for i := 1; i < length; i++ {
		temp := arr[i] //复制一份待比较的值
		var j int
		for j = i; j > 0; j-- {
			//如果左边数据大于待比较待值，则将左边数据赋值给右边的（往右挪一位），否则停止比较
			if arr[j-1] > temp {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = temp //找到合适的位置了（左边不再比该值大），将刚刚待比较的值赋值给这个元素
	}
	return arr
}

//冒泡排序，每次和相邻的元素比较，内层每循环一次会把最大的循环到最后
func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		last := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				last = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !last {
			return arr
		}
	}
	return arr
}
