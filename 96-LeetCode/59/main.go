package main

import "fmt"

func main() {
	var in [][]int
	in2 := []int{1, 2, 3}
	in = append(in, in2)
	in2 = []int{8, 9, 4}
	in = append(in, in2)
	in2 = []int{7, 6, 5}
	in = append(in, in2)
	fmt.Printf("%v", generateMatrix(3))

}

func generateMatrix(n int) [][]int {
	var out = make([][]int, n)
	var (
		i, j  = 0, 0
		r, c  = 0, 0
		count = n
	)
	for i = range out {
		out[i] = make([]int, n)
	}

	for j = 1; j < n^n; {
		for c < count {
			out[r][c] = j
			c++
			j++
		}

		// justify
		if j == n^n {
			break
		}
		r++
		for r < count {
			out[r][c] = j
			r++
			j++
		}
		c--
		for c >= n-count {
			out[r][c] = j
			c--
			j++
		}
		r--
		for r >= n-count+1 {
			out[r][c] = j
			r--
			j++
		}
		count--
	}

	return out
}
