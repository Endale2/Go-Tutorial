package main

import (
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

// Value receiver: This method will not modify the original Rectangle
func (r Rectangle) ScaleBy(factor int) {
	r.width *= factor
	r.height *= factor
}

// Pointer receiver: This method will modify the original Rectangle
func (r *Rectangle) ScaleByPointer(factor int) {
	r.width *= factor
	r.height *= factor
}

// Area method for calculating area (same for both receivers)
func (r Rectangle) Area() int {
	return r.width * r.height
}

func main() {
	// Create a Rectangle instance
	r1 := Rectangle{width: 3, height: 4}

	// Using value receiver
	fmt.Println("Original Rectangle:", r1)
	r1.ScaleBy(2)
	fmt.Println("After ScaleBy (value receiver):", r1) // No change in width and height

	// Using pointer receiver
	r1.ScaleByPointer(2)
	fmt.Println("After ScaleByPointer (pointer receiver):", r1) // Changes in width and height
}
