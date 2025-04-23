package main

import "fmt"

func main() {

	fmt.Println("one main start")
	defer myDefer()
	 myDefer2()
	fmt.Println("two main end")
	fmt.Println("three main end")
}

func myDefer() {
	fmt.Println("------myDefer start-----")
	defer func() {
		fmt.Println("---myDefer end 01---")
	}()
	fmt.Println("---- myDefer end -----")
}
func myDefer2() {
	fmt.Println("----myDefer2 start -----")
	defer func() {
		fmt.Println("myDefer2 end------")
	}()

	fmt.Println("---- myDefer2 end -----")
}
