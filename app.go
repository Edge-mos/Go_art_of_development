package main

import "fmt"

func main() {
	fmt.Println("Defer")

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	fmt.Println("defer is stack - LIFO")
}
