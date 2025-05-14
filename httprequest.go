package main

import (
	"fmt"
	"net/http"
)
const url="http://www.baidu.com";
func main() {
	fmt.Println("welcome to the web request of Go!")
	getRequest()
	fmt.Println("welcome to the web request of Go!")
}



func getRequest() {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
		return
	}
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Headers: %v\n", resp.Header)
	fmt.Printf("Response Body: %v\n", resp.Body)
	fmt.Printf("Response Content Length: %d\n", resp.ContentLength)
	fmt.Printf("Response Request: %T\n", resp)
	defer resp.Body.Close()

}