package main

import "fmt"

// Interface
type Shape interface {
	Area() float64
}

// Structs
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

// Method to implement Shape interface for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Method to implement Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Function to calculate total area of shapes
func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	// Polymorphism: Different types implementing the same interface
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 3, Height: 4},
	}

	// Pass by value: Original values are not modified
	circle := Circle{Radius: 10}
	modifyValue(circle)
	fmt.Println("Radius after modifyValue:", circle.Radius) // Output: 10 (original value remains unchanged)

	// Pass by reference: Original values can be modified indirectly
	modifyPointer(&circle)
	fmt.Println("Radius after modifyPointer:", circle.Radius) // Output: 20 (original value is modified)

	// Interface usage: Total area of all shapes
	fmt.Println("Total area of shapes:", TotalArea(shapes)) // Output: Total area of shapes: 91
}

// Function to modify value (pass by value)
func modifyValue(c Circle) {
	c.Radius = 20 // This modification only affects the local copy of c
}

// Function to modify value indirectly via pointer (pass by reference)
func modifyPointer(c *Circle) {
	c.Radius = 20 // This modification affects the original variable
}
