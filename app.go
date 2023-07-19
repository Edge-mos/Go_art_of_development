package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go_if_statements")

	for i := 0; i < 1000; i++ {
		if i%100 == 0 {
			fmt.Printf("i = %d \n", i)
		}
	}

	fmt.Println("-------------------------")

	var init = 0

	for ; init < 1000; init++ {
		if init%100 == 0 {
			fmt.Printf("init = %d \n", init)
		}
	}

	fmt.Println("-------------------------")

	var b int16 = 0

	for b < 1000 {
		if b%100 == 0 {
			fmt.Println(fmt.Sprintf("event! b = %d", b))
		}
		b++
	}

	fmt.Println("-------------------------")

	var counter int16 = 0
	for {
		if counter%100 == 0 {
			fmt.Println(fmt.Sprintf("Warn! counter = %d", counter))
		}

		if counter == 1000 {
			fmt.Println("Exit!")
			break
		}
		counter++
	}

	fmt.Println("-------------------------")

	for val := 0; val <= 1000; val++ {

		if result := isEvent(val); result == true { // вот тут происходит инициализация с присваиванием в result и последующее условие
			fmt.Println(fmt.Sprintf("result is true! val = %d", val))
		}

	}

	fmt.Println("------------------------- Switch")

	for c := 0; c < 5; c++ {
		switch c {
		case 1:
			fmt.Println("One!")
		case 2:
			fmt.Println("Two!")
		case 3:
			fmt.Println("THREE!")
		default:
			fmt.Println("Wat?")
		}
	}

}

func isEvent(val int) bool {
	if val%100 == 0 {
		return true
	}
	return false
}
