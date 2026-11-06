package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	x float64
	y float64
}

func (r Rect) Area() float64 {
	return r.x * r.y
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.x + r.y)
}

func (r Rect) String() string {
	return fmt.Sprintf("Rect(x=%.2f , y=%.2f)", r.x, r.y)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(r=%.2f)", c.radius)
}

type Triangle struct {
	a, b, c float64
}

func (t Triangle) Area() float64 {
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle(a=%.2f,b=%.2f,c=%.2f)", t.a, t.b, t.c)
}

func Largest(shapes []Shape) Shape {
	if len(shapes) == 0 {
		return nil
	}
	var largestShape Shape = shapes[0]

	for _, shape := range shapes {
		if shape.Area() > largestShape.Area() {
			largestShape = shape
		}
	}
	return largestShape
}

func main() {

	shapes := []Shape{
		Rect{3, 4},
		Rect{10, 2},
		Circle{1},
		Circle{2},
		Triangle{3, 4, 5},
		Triangle{5, 5, 5},
		Triangle{6, 6, 10},
	}

	fmt.Println(Largest(shapes))
}
