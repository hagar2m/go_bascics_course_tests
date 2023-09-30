package main

import (
	"fmt"
)

// Define an interface named "Shape" with a method called "Area".
type Shape interface {
	Area() float64
}

// Define two struct types: Circle and Rectangle.

// Circle represents a circle with a radius.
type Circle struct {
	Radius float64
	Name   string
}

// Implement the Area method for Circle.
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Rectangle represents a rectangle with length and width.
type Rectangle struct {
	Length float64
	Width  float64
	Name   string
}

// Implement the Area method for Rectangle.
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func main() {
	// Create instances of Circle and Rectangle.
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Length: 4, Width: 6}

	// Use the Area method through the Shape interface.
	shapes := []Shape{circle, rectangle}

	// Calculate and print the areas of the shapes.
	for _, shape := range shapes {
		fmt.Printf("Area of shape: %.2f\n", shape.Area())
	}
}
