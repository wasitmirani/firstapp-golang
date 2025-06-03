package main

import (
	"fmt"
	"strings"
)

type Course  struct{
	Name  string
	Price float32
	Provider string
	ID string
	Tags  []string

}

func main() {
	fmt.Println("welcome to go lang json format")
	EncodeJson()
}


func EncodeJson(){


	courseList := []Course{
		{"Test",10.2,"AWS","123",[]string{"1","2","Test"}},
{
			Name:     "Advanced Go Programming",
			Price:    49.99,
			Provider: "Google",
			ID:       "456",
			Tags:     []string{"Go", "Programming", "Advanced"},
		},
		{
			Name:     "Cloud Computing Fundamentals",
			Price:    29.99,
			Provider: "Azure",
			ID:       "789",
			Tags:     []string{"Cloud", "Beginner", "Fundamentals"},
		},
	}

	
		for i,course := range courseList{
	     fmt.Printf("%d. %s (ID: %s)\n", i+1, course.Name, course.ID)
		fmt.Printf("   Provider: %s\n", course.Provider)
		fmt.Printf("   Price: $%.2f\n", course.Price)
		fmt.Printf("   Tags: %s\n\n", strings.Join(course.Tags, ", "))
		}

}




