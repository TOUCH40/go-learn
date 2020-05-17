package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	fmt.Println(len(myGreeting))
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
	fmt.Println(len(myGreeting))
}
