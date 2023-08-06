package main

import (
	"fmt"
)

func main() {
	fmt.Println("Interfaces")

	var a InterfaceHere

	structHere := StructHere{N1: 1, N2: 2}
	fmt.Printf("structHere.Sum() = %d \n", structHere.Sum())

	otherStruct := OtherStruct{A: 2, B: 3}
	fmt.Printf("otherStruct.Sum() = %d \n", otherStruct.Sum())

	a = &structHere
	fmt.Printf("a.Sum() = %d \n", a.Sum())

	a = otherStruct
	fmt.Printf("a.Sum() other = %d \n", a.Sum())

	var i InterfaceHere = OtherStruct{A: 3, B: 3}
	fmt.Printf("(%v , %T)\n", i, i)

	//empty interface
	var emp interface{}
	emp = "jelly"
	fmt.Printf("(%v , %T)\n", emp, emp)
	emp = 42
	fmt.Printf("(%v , %T)\n", emp, emp)

	//type cast
	emp = "hello"
	s, ok := emp.(string)
	fmt.Println(s, ok)

	f, ok := emp.(float32)
	fmt.Println(f, ok)

}

// Interface_section -------------------------------

type InterfaceHere interface {
	Sum() int
}

// Struct_section-----------------------------------

type StructHere struct {
	N1, N2 int
}

func (sh *StructHere) Sum() int {
	return sh.N1 + sh.N2
}

type OtherStruct struct {
	A, B int
}

func (os OtherStruct) Sum() int {
	return os.A + os.B
}
