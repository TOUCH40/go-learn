package main

import "fmt"

func main() {
	In := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(In))
}

func dailyTemperatures(T []int) []int {
	Length := len(T)
	Out := make([]int, Length)
	for i := 0; i < Length; i++ {
		for j := i; j < Length; j++ {
			if T[j] > T[i] {
				Out[i] = j - i
				break
			}
			if j == Length-1 {
				Out[i] = 0
			}
		}
	}
	return Out
}
