package main

import "fmt"

func main() {
	in := make([][]int, 2)
	for i := 0; i < 2; i++ {
		in[i] = make([]int, 2)
	}
	in[0][0] = 1
	in[0][1] = 2
	in[1][0] = 1
	in[1][1] = 3
	fmt.Println(kthSmallest(in, 1))
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	if k == 1 {
		return matrix[0][0]
	}
	r := 0
	c := 0
	for i := 0; i <= k; i++ {
if {

}	else	if matrix[r][c+1] > matrix[r+1][c] {
			c = c + 1
		} else {
			r = r + 1
		}
	}
	// r := k / n
	// c := k%n - 1
	return matrix[r][c]
}
