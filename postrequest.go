package main


import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main(){
    fmt.Println("Hello, World!")
}



func postRequest() {
	// Create a new HTTP POST request
	resp, err := http.Post("http://example.com", "application/json", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the response status and body
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}