package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := []int{3, 43, 48, 94, 85, 33, 64, 32, 63, 66}

	fmt.Println(minNumber(n))
}

func minNumber(nums []int) string {
	words := []string{}
	for _, item := range nums {
		words = append(words, strconv.Itoa(item))
	}
	quickSort(words, 0, len(nums)-1)
	return strings.Join(words, "")
}

func quickSort(a []string, l int, r int) {
	// 跳出
	if l >= r {
		return
	}
	// 基准，左，右
	pivot, ol, or := a[l], l, r
	for ol < or {
		for ol < or && (compare(a[or], pivot) == 0 || compare(a[or], pivot) == 1) {
			or--
		}
		a[ol] = a[or]
		for ol < or && compare(pivot, a[ol]) == 1 {
			ol++
		}
		a[or] = a[ol]
	}
	a[ol] = pivot
	quickSort(a, l, ol-1)
	quickSort(a, ol+1, r)
}

func compare(a string, b string) int {
	i1, _ := strconv.Atoi(a + b)
	i2, _ := strconv.Atoi(b + a)
	if i1 > i2 {
		return 1
	} else if i1 < i2 {
		return -1
	} else {
		return 0
	}
}

// func quickSort(a []int, l int, r int) {
// 	// 跳出
// 	if l >= r {
// 		return
// 	}
// 	// 基准，左，右
// 	pivot, ol, or := a[l], l, r
// 	for ol < or {
// 		for ol < or && a[or] >= pivot {
// 			or--
// 		}
// 		a[ol] = a[or]
// 		for ol < or && a[ol] < pivot {
// 			ol++
// 		}
// 		a[or] = a[ol]
// 	}
// 	a[ol] = pivot
// 	quickSort(a, l, ol-1)
// 	quickSort(a, ol+1, r)
// }

// 快速排序实现，由于要二分递归，所以要传入每次排序的起始和结束的索引
func quick_sorting(data []int, start, end int) {
	if start < end {
		// 获取基准值
		base := data[start]
		// 临时变量，左右索引
		left := start
		right := end
		// 进入循环
		for left < right {
			// 由于左边的(第0个)被取走当成基准值，所以在左边就留下一个洞，用于存储比基准小的值
			// 所以先从右边找，找到第一个比基准值小的
			for left < right && data[right] >= base {
				right--
			}
			// 找到了比基准值小的，那就把这个值填入左边的洞，做索引要++
			if left < right {
				data[left] = data[right]
				left++
			}
			// 因为上面的操作让右边留了一个洞，开始从左边找比基准值大的
			for left < right && data[left] <= base {
				left++
			}
			// 找到比基准值大的再次把右边洞填上，并且在左边又留了一个洞
			if left < right {
				data[right] = data[left]
				right--
			}
		}

		// 把基准值填入到洞中，这就是本应该他所在的位置
		data[left] = base
		// 继续分两组递归排序，注意此时基准值已经不用参与排序了
		quick_sorting(data, start, left-1)
		quick_sorting(data, left+1, end)
	}
}

// func quickSort(arr []int, start, end int) {
// 	if start < end {
// 		i, j := start, end
// 		key := arr[(start+end)/2]
// 		for i <= j {
// 			for arr[i] < key {
// 				i++
// 			}
// 			for arr[j] > key {
// 				j--
// 			}
// 			if i <= j {
// 				arr[i], arr[j] = arr[j], arr[i]
// 				i++
// 				j--
// 			}
// 		}

// 		if start < j {
// 			quickSort(arr, start, j)
// 		}
// 		if end > i {
// 			quickSort(arr, i, end)
// 		}
// 	}
// }
