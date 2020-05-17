package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}

func main() {
	fmt.Println(makeGreeter()())
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
}