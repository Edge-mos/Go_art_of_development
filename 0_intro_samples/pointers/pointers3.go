package main

import "fmt"

func main() {

	p := new(int)
	fmt.Println(*p) // будет 0, так как int по-умолчанию
	//-----------------------------------------------------

	i := newInt()
	i2 := newIntV2()

	fmt.Println("i и i2", *i, *i2)

	//var tmp *int
	//fmt.Println("!!! тут паника, ссылка никуда не указывает - nil. нельзя разыменовать nil", *tmp)
}

/*
 * создает неименованную переменную и инициализирует ее нулевым значение
 */
func newInt() *int {
	return new(int)
}

/*
 * тоже самое
 */
func newIntV2() *int {
	var dummy int
	return &dummy
}
