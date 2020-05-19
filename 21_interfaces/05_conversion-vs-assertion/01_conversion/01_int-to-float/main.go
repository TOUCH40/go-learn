package main

import "fmt"

func main() {
	var x = 12
	var y = 12.1230123
	fmt.Println(y + float64(x))
	// conversion: int to float64
}

/*my-test*/
// func main() {
// 	var iv int32
// 	var fv float64 = 89.09
// 	iv = int32(fv)
// 	fmt.Println(fv)
// 	fmt.Println(iv)
// }
