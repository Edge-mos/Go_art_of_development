package main

import "fmt"

func main() {
	x := 1
	p := &x

	fmt.Println("Показывает адрес переменной: ", p)
	fmt.Println("А это разъименование переменной: ", *p)

	*p = 2
	fmt.Println("Теперь изначальная переменная х = ", x)
	//----------------------------------------------------------------------------------------
	fmt.Println("два указателя равны только в том случае, если они указывают на одну переменную или же оба равны nil")
	var pntOne *int
	var pntTwo *int

	fmt.Println("указатели равны? : ", pntOne == pntTwo)

	pntOne = &x
	pntTwo = p

	fmt.Println("а теперь указатели равны? : ", pntOne == pntTwo)
	//----------------------------------------------------------------------------------------

	res := retPointer()
	fmt.Println("retPointer func call: ", res)
	res2 := retPointer()
	fmt.Println("retPointer func call2: ", res2)

}

func retPointer() *int {
	val := 1
	return &val
}
