package main

import (
	"fmt"
	"os"
)

//Проходка по входящим аргументам v2

func main() {
	args := os.Args
	accumulator, separator := "", "_"

	for _, arg := range args[1:] {
		accumulator += arg + separator
	}

	fmt.Printf("Result echo2 = %v", accumulator)
}
