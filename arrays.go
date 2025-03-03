package main

import "fmt"

func main() {
	// Declare an array of integers
	var intArray [5]int
	intArray[0] = 10
	intArray[1] = 20
	intArray[2] = 30
	intArray[3] = 40
	intArray[4] = 50

	// Declare and initialize an array of strings
	stringArray := []string{"Go", "Python", "Java"}

	// Print the arrays
	fmt.Println("Integer Array:", intArray)
	fmt.Println("String Array:", stringArray)
}