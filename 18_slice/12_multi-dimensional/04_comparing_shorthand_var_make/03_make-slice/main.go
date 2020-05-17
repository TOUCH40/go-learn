package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "Todd"
	fmt.Println(cap(student))
	student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(cap(student))
}
