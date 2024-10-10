package main

import "fmt"

func main() {
	// Declare and initialize an array
	arr := [4]int{10, 20, 30, 40}

	// Access elements
	fmt.Println("First element:", arr[0])

	// Modify an element
	arr[1] = 25
	fmt.Println("Modified array:", arr)

	// Iterate over the array using a range loop
	fmt.Println("Array elements:")
	for index, value := range arr {
		fmt.Printf("Index %d: %d\n", index, value)
	}

	// Length of the array
	fmt.Println("Length of the array:", len(arr))

	// Multidimensional array (2D array)
	matrix := [2][2]int{
		{1, 2},
		{3, 4},
	}

	fmt.Println("2D array (matrix):")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
