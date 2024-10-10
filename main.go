package main

import "fmt"

func main() {
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("It's the start of the week.")
	case "Tuesday":
		fmt.Println("It's Tuesday!")
	case "Friday":
		fmt.Println("It's almost the weekend!")
	default:
		fmt.Println("It's a regular day.")
	}
}
