package main

import (
	"fmt"
)

func main() {

	r1 := Rectangle{3, 4}

	fmt.Printf("The area of the rectangle is : %v ", r1.Area())

}

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) Area() int {
	return r.height * r.width
}
