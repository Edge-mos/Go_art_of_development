package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	var variable = "Hello World"
	var pointer = &variable

	fmt.Println(variable)                              // print value
	fmt.Printf("Pointer : %v \n", pointer)             // print pointer in hex-format (16)
	fmt.Printf("Pointer dereference: %v \n", *pointer) // dereference - get value thru pointer

	*pointer = "oh my!"
	fmt.Printf("New value thru pointer : %v \n", variable)

}
