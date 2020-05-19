package main

import "fmt"

func main() {
	var x rune = 'a' // rune is an alias for int32; normally omitted in this statement
	var y int32 = 'b'
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(string(x))
	fmt.Println(string(y))
	fmt.Printf("%T", string(x))
	// conversion: rune to string
	// fmt.Println(y.(string))
}
