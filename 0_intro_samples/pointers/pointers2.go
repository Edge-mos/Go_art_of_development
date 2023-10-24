package main

import "fmt"

func main() {
	//---------------------------------------------
	x := 1
	increment(x)
	fmt.Println("after increment func = ", x) //не мутирует - в функцию попадает по значению
	//---------------------------------------------
	i := 1
	fmt.Println("ссылка до функции", &i)
	incrementPnt(&i)
	fmt.Println("after incrementPnt func = ", i) //мутирует - в функцию попадает по ссылке(можно сравнить в функции)
	//---------------------------------------------
	j := 11
	fmt.Println(&j)
	incrementPntV2(j)
	fmt.Println("after incrementPntV2 func = ", j) //не мутирует - в функцию попадает по значению, а только операции с указателем
	//---------------------------------------------
}

func increment(val int) {
	val++
}

func incrementPnt(val *int) {
	fmt.Println("внутри функции: ", val)
	*val++
}

func incrementPntV2(val int) {
	fmt.Println("внутри функции: ", &val)
	*&val++
}
