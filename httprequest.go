package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const endpoint="http://www.baidu.com";
func main() {
	fmt.Println("welcome to the web request of Go!")
	// getRequest()
	fmt.Println("welcome to the web request of Go!")
	getURLData()
	
}


func getURLData(){

	result, err := url.Parse("http://www.baidu.com?query=golang&name=zhangsan#fragment")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
		return
	}
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Raw Path:", result.RawPath)
	fmt.Println("Raw Query:", result.RawQuery)
	fmt.Println("Raw Fragment:", result.RawFragment)

	fmt.Printf("Raw Port:", result.Port())
	fmt.Println("Raw User:", result.User)
	fmt.Println("Raw User Info:", result.User)

	queryparanms := result.Query()
	fmt.Printf("Query Parameters: %s", queryparanms["name"])

}

func getRequest() {
	resp, err := http.Get(endpoint)

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
	databytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
		return
	}

	content := string(databytes)
	fmt.Println("Response Content:", content)

}