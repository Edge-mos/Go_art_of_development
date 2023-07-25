package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var arr [2]string
	arr[0] = "hello"
	arr[1] = "world"

	fmt.Println(arr)

	var numbers = [...]int{1, 2, 3}

	fmt.Println(numbers)
}
