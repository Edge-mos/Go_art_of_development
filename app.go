package main

import "fmt"

func main() {
	fmt.Println("Structures")

	var point = Point{
		X: 1,
		Y: 2,
		S: "Some_message",
	}

	fmt.Printf("point : %v \n", point)

	var point2 = Point{X: 123}

	fmt.Printf("point2 : %v \n", point2)

	var p = &point

	fmt.Printf("pointer of struct: %v \n", p)
	fmt.Printf("pointer of struct property: %v \n", p.X)
	fmt.Printf("dereference of pointer: %v \n", *p)

	point.toString()

}

type Point struct {
	X int
	Y int
	S string
}

func (point Point) toString() {
	fmt.Printf("X: %v \n", point.X)
	fmt.Printf("Y: %v \n", point.Y)
	fmt.Printf("S: %v \n", point.S)
}
