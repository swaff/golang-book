package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimiter() float64
}

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) perimiter() float64 {
	return math.Pi * c.radius * 2
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

func (rec *Rectangle) area() float64 {
	length := rec.x2 - rec.x1
	height := rec.y2 - rec.y1
	return length * height
}

func (rec *Rectangle) perimiter() float64 {
	length := rec.x2 - rec.x1
	height := rec.y2 - rec.y1
	return 2 * (length + height)
}

func main() {
	circle := Circle{radius: 5}
	fmt.Println("Circle area is:", circle.area())
	fmt.Println("Circle perimiter is:", circle.perimiter())

	rectangle := Rectangle{x1: 1, x2: 100, y1: 1, y2: 50}
	fmt.Println("Rectangle area is:", rectangle.area())
	fmt.Println("Rectangle perimiter is:", rectangle.perimiter())
}
