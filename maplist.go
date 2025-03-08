package main

import "fmt"

func main() {

	languages := make(map[string]string)
	languages["Php"] = "Laravel"
	languages["Java"] = "Spring"
	languages["Go"] = "Gin"
	languages["Python"] = "Django"
	languages["Ruby"] = "Rails"
	languages["JavaScript"] = "NodeJs"
	languages["C#"] = "ASP.NET"
	languages["C++"] = "Qt"
	languages["C"] = "Linux Kernel"
	fmt.Println("list of languages fremworks: ",languages)

	// delete a key from the map
	delete(languages, "C++")
	fmt.Println("list of languages fremworks after deleting C++: ",languages)


	for key, value := range languages {
		fmt.Println(key, ":", value)
	}

	// check if a key exists in the map
	_, ok := languages["C++"]
	if ok {
		fmt.Println("C++ exists in the map")
	} else {
		fmt.Println("C++ does not exist in the map")
	}


}