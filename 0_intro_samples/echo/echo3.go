package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	res := strings.Join(os.Args[1:], ", ")

	fmt.Printf("Res: %v", res)
}
