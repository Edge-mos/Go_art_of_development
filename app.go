package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println("Reader interface")

	reader := strings.NewReader("Hello World")
	sls := make([]byte, 8)

	fmt.Printf("initial slice: %v \n", sls)

	for {
		n, err := reader.Read(sls)
		fmt.Printf("n = %d err = %v sls = %v ", n, err, sls[:n])
		fmt.Printf("sls --> %q \n", sls[:n])

		if err == io.EOF {
			break
		}
	}

	fmt.Println(sls)

}
