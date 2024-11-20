package main

import "fmt"

func main() {

	result := sub(9, 4)
	fmt.Print(result)
}

func sub(x int, y int) int {
	return x - y
}
