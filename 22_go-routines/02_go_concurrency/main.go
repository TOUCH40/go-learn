package main

import (
	"fmt"
)

func main() {
	go foo()
	go bar()
	var cd string
	fmt.Scanln(cd)
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
