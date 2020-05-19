package main

import (
	"fmt"
	"math"
)

type shape interface {
	calPerimeter() float64
}

type circle struct {
	Radius float64
}

func (c circle) calPerimeter() float64 {
	return c.Radius * c.Radius * math.Pi
}

type squire struct {
	SideLength float64
}

func (s squire) calPerimeter() float64 {
	return s.SideLength * 4
}

func getLength(s shape) {
	fmt.Println(s.calPerimeter())
}

func main() {
	c := circle{10}
	s := squire{10}
	getLength(c)
	getLength(s)
}
