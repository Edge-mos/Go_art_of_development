package main

import (
	"art_of_development/0_intro_samples/packages_examples/conv"
	"fmt"
)

func main() {
	var celsius conv.Celsius = 101
	fmt.Println(celsius)

	tmp := conv.Celsius(12)
	fmt.Println(tmp)

	zero := conv.AbsoluteZeroC
	fmt.Println("Absolute zero: ", zero)
}
