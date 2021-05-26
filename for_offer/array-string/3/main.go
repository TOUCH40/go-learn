package main

func main() {
	in := []int{}
}

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := (left + right) >> 1
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[left] {
			left = mid
		} else {
			right--
		}
	}
	return numbers[left]
}
