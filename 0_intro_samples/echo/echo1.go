package main

// Проходка по входящим аргументам

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	var separator = " "
	var accumulator string

	for i := 0; i < len(args); i++ {
		accumulator += args[i] + separator
	}

	fmt.Printf("Result: %v", accumulator)

}
