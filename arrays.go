package main

import (
	"fmt"
	"sort"
)

func main() {
	// Declare an array of integers
	var intArray [5]int
	intArray[0] = 10
	intArray[1] = 20
	intArray[2] = 30
	intArray[3] = 0
	intArray[4] = 50

	// Declare and initialize an array of strings
	stringArray := []string{"Go", "Python", "Java"}

	// Print the arrays
	fmt.Println("Integer Array:", intArray)
	fmt.Println("String Array:", stringArray)
	sliceArray()
}

func removeElement(array []string, index int) []string {
	return append(array[:index], array[index+1:]...)
}
func sliceArray( ){
	// Declare and initialize a slice
	slice := []int{10, 20, 30, 40, 50}
	// Print the slice
	fmt.Println("Slice:", slice)

	// Print the length of the slice
	fmt.Println("Length:", len(slice))
	fmt.Printf("Length: %T\n", slice)
	// Print the capacity of the slice
	fmt.Println("Capacity:", cap(slice))
	fruitList := []string{"Apple", "Banana", "Mango"}
	fruitList = append(fruitList, "Orange");
	fruitList = append(fruitList, "Apple2")
	fruitList = append(fruitList, "Banana2")
	fruitList = append(fruitList, "Mango2")
		
	fruitList= removeElement(fruitList, 2)

	fmt.Println(fruitList)


	// Declare and initialize a slice
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	// Print the slice
	fmt.Println("High Scores:", highScores)


	highScores = append(highScores, 55,666,99,1000)
	sort.Ints(highScores)
	fmt.Println("High Scores:", highScores)
	fmt.Println("Length:", sort.IntsAreSorted(highScores);
	

)

}