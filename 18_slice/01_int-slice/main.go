package main

import "fmt"

func main() {

	mySlice := []int{1, 3, 5, 7, 9, 11}
	change(mySlice)
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
}
func change(param []int) {
	param[0] = 10
}
