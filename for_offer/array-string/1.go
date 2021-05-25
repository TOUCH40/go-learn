package main

import "fmt"

func main() {
	// arr := [][]int{
	// 	{1, 4, 7, 11, 15},
	// 	{2, 5, 8, 12, 19},
	// 	{3, 6, 9, 16, 22},
	// 	{10, 13, 14, 17, 24},
	// 	{18, 21, 23, 26, 30}}
	arr := [][]int{{-5}}
	fmt.Println(findNumberIn2DArray(arr, -5))

}
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	for i, j := m-1, 0; i >= 0 && j < n; {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}

	}
	return false
}

// 暴力算法
// func findNumberIn2DArray(matrix [][]int, target int) bool {
// 	if matrix == nil || len(matrix) == 0 {
// 		return false
// 	}
// 	m, n := len(matrix), len(matrix[0])
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if matrix[i][j] == target {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
