package main

import "fmt"

func main() {
	
	myNumber :=10;
	var myPointer = &myNumber
	fmt.Println("Value of myNumber is: ", myNumber)
	fmt.Println("Address of myNumber is: ", &myNumber)
	fmt.Println("Value of myPointer is: ", myPointer)
	fmt.Println("Value of myPointer is: ", *myPointer)
	
	*myPointer = *myPointer * 2
	fmt.Println("Value of myNumber is: ", myNumber)
	fmt.Println("Value of myPointer is: ", *myPointer)
}