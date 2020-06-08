package main

import "fmt"

func main() {
	matrix := [][]int{}
	ret := spiralOrder(matrix)
	fmt.Println(ret)
}

func spiralOrder(matrix [][]int) []int {
	out := []int{}
	row := len(matrix)
	if row == 0 {
		return out
	}
	col := len(matrix[0])
	c0 := 0
	c1 := col - 1
	r0 := 0
	r1 := row - 1
	for c0 <= c1 && r0 <= r1 {

		for i := c0; i <= c1; i++ {
			out = append(out, matrix[r0][i])
		}
		for i := r0 + 1; i <= r1; i++ {
			out = append(out, matrix[i][c1])
		}
		if r1 > r0 {
			for i := c1 - 1; i >= c0; i-- {
				out = append(out, matrix[r1][i])
			}
		}
		if c1 > c0 {
			for i := r1 - 1; i >= r0+1; i-- {
				out = append(out, matrix[i][c0])
			}
		}

		c0++
		c1--
		r0++
		r1--
	}

	return out
}
