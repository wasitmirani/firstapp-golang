package main

import "fmt"

func main() {

	fmt.Println("welcome to loop in golang")
	days := []string{"mon", "sun", "sat"}

	fmt.Println(days)
	for i := 0; i < len(days); i++ {
		fmt.Printf("value is %v\n", days[i])
	}


	// Using range loop like foreach loops
	for index, day := range days{
		fmt.Printf("value is %v at index %v\n", day, index)
		fmt.Println()
	}

	// Using for loop with a condition
	for i := 0; i < 5; i++ {
        fmt.Printf("value is %v\n", i)
    }

    // Using for loop with a break statement
    for i := 0; i < 5; i++ {
        fmt.Printf("value is %v\n", i)
        if i == 2 {
            break
        }
    }

    // Using for loop with a continue statement
    for i := 0; i < 5; i++ {
        if i == 2 {
            continue
        }
        fmt.Printf("value is %v\n", i)
    }
    fmt.Println()

    // Using for loop with a range statement
    for i, day := range days {
        fmt.Printf("value is %v at index %v\n", day, i);
    }

	val :=1;
	
	for val < 10 {
		val++
		if( val == 2){
			continue;
		}
		if val == 4{
			goto jump;
		}
		if( val == 5){
			break;
		}
        fmt.Printf("array value is: %v\n", val)
        
    }
	jump: fmt.Println("Jumped to label")
}
