package main

import "fmt"

func main() {
	A := []int{1, 2, 3, 2, 1}
	B := []int{3, 2, 1, 4, 7}
	fmt.Println(findLength(A, B))
}

func findLength(A []int, B []int) int {
	out := 0
	n := len(A)
	m := len(B)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := n - 1; i > -1; i-- {
		for j := m - 1; j > -1; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if out < dp[i][j] {
				out = dp[i][j]
			}
		}
	}

	return out
}
