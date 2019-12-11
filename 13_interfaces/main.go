package main

import (
	"fmt"
	"math"
)

// define interface
type Shape interface {
	// method
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Reactangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Reactangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Reactangle{width: 10, height: 5}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))

}
