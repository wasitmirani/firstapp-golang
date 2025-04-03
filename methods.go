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
	wasit := User{"John",10,"wasit@example.com","test abc street address","","",nil,true};
	fmt.Println(wasit.GetStatus());
	fmt.Println("Result--------------------------------")
	wasit.PrintDetails();
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


type User struct {
	Name string
	Age  int
	Email string
	Address string
	Phone string
	Gender string
	Hobbies []string
	Status bool
}


func (u *User) PrintDetails() {
	fmt.Printf("Name: %s\nAge: %d\nEmail: %s\nAddress: %s\nPhone: %s\nGender: %s\nHobbies: %v\nStatus: %t\n", u.Name, u.Age, u.Email, u.Address, u.Phone, u.Gender, u.Hobbies, u.Status)
}

func (u *User) UpdateDetails(name string, age int, email string, address string, phone string, gender string, hobbies []string, status bool) {
	u.Name = name
    u.Age = age
    u.Email = email
    u.Address = address
    u.Phone = phone
    u.Gender = gender
    u.Hobbies = hobbies
    u.Status = status
}

func (u *User) ValidateEmail() bool {
    // Add email validation logic here
    return true // Placeholder, replace with actual validation logic
}

func (u *User) GetStatus() string{
	return fmt.Sprintf("User is online :%v", u.Status);
}