package main

import "fmt"

func main() {
	result := add(3, 4)
	fmt.Println("result:  ", result)
}

func add(a int, b int) int {
	return a + b
}
