package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("welcome to golang fuunctions")
	fmt.Println(sum(5, 3))
	fmt.Println(multiply(5, 3))
	fmt.Println(getNamed("John Doe"))
	fmt.Println(greet("John Doe"))
	fmt.Println(addValues(5, 3, 2, 1))
	fmt.Println(addNames("John", "Doe", "Jane"))
	result,response := getNumAndMessage(1,5,4)
	fmt.Printf("Result: %d, Message: %s\n", result, response)
}

func sum(a int, b int) int {
    return a + b
}

func multiply(a int, b int) int {
    return a * b
}


func getNamed(name string) string{

	return name;
}

func greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}


func addValues(values ...int) int{
	sum := 0
    for _, value := range values {
        sum += value
    }
    return sum
}

func addNames(names ...string) string{
	return fmt.Sprintf("Hello, %s!", strings.Join(names, ", "))
}

func  getNumAndMessage(num ...int) (int, string) {
	if len(num) == 0 {
        return 0, "No numbers provided"
    }
    return num[0], fmt.Sprintf("The number is: %d", num[0])
}