package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main(){
    fmt.Println("Hello, World!")
	performPostRequest()
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

func performPostRequest(){
	print("hello")
	const myurl ="https://oms.zip24.com/dashboard/overview";

	//form data
	data := url.Values{};
	data.Add("firstname","wasit")
	data.Add("lastname","mirani")
	data.Add("email","miraniwasit2@gmail.com")
	println("this request for post data")
	res, err := http.PostForm(myurl,data)
	if err !=nil{

		panic(err)
	}


	defer res.Body.Close()

	constant, _ := io.ReadAll(res.Body)

	fmt.Println(string(constant))

}