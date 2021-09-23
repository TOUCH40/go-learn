package main

import (
	"fmt"
	"strings"
)

var PI float32 = 3.14

type Shape interface {
	Perimeter() float32
	Area() float32
}

type Circle struct {
	// radius float32
	Diameter float32
}

func (c *Circle) SetDiameter(v float32) {
	c.Diameter = v
}

func (c *Circle) Perimeter() float32 {
	return 2 * (c.Diameter / 2) * PI
}

func (c *Circle) Area() float32 {
	return (c.Diameter / 2) * (c.Diameter / 2) * PI
}

func main() {
	// var ci Shape = &Circle{Diameter: 10}
	// ci2 := ci.(*Circle)
	// ci2.SetDiameter(20)
	// println(ci2.Perimeter())
	test := strings.Builder{}
	// res := fmt.Sprintf("go %f", ci.Perimeter())
	// println(res)
	test.WriteString("fame of pain")
	test.WriteString("fame of pain")
	fmt.Println(test.String())
	test_recover()

}

func test_recover() {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {
			fmt.Println("recover success")
		}
	}()
	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
	fmt.Println("after panic")
}
