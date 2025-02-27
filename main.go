package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)


func main()  {
	maxProcs := runtime.GOMAXPROCS(0)
    numCPU := runtime.NumCPU()
	print("Number of CPUs: ", numCPU)
	print("Max Procs: ", maxProcs)
	welcome := "welcome to the world of Go"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello %T", name)

}