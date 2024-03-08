package main

import (
	"fmt"
	"math"
)

func main() {

	rect := Rectangle{30, 40}
	fmt.Println(getArea(rect))

	circ := Circle{5}
	fmt.Println(getArea(circ))
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (rect Rectangle) area() float64 {

	return rect.width * rect.height
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}
