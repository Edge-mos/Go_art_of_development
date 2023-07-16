package main

import "fmt"

func main() {

	//standard
	var sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//while analog
	var sum2 = 0
	for sum2 < 1000 {
		sum2 += 10
	}

	fmt.Println(sum2)

}
