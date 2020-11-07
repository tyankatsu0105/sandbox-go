package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{width: 12, height: 2}
	r2 := Rectangle{width: 9, height: 4}
	c1 := Circle{radius: 10}
	c2 := Circle{radius: 25}

	fmt.Println("r1", r1.area())
	fmt.Println("r2", r2.area())
	fmt.Println("c1", c1.area())
	fmt.Println("c2", c2.area())
}
