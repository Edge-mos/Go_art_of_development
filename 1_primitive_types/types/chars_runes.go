package types

import (
	"fmt"
)

func CharsRunesEx() {
	fmt.Println("=====руны, форматирование========")
	//руны, символы выводятся с помощью %[1]c или каст string() или %[1]q, если в кавычках

	ascii := 'a'
	fmt.Printf("%d ; %[1]c ; %[1]q \n", ascii)
	fmt.Printf("каст до стринг %s \n", string(ascii))

	var tmp rune = 'a'
	fmt.Printf("каст до стринг %s \n", string(tmp))

	//var imgPlane rune = '✈'
	var imgPlane rune = 9992
	fmt.Println(string(imgPlane))
	fmt.Printf("%c\n", imgPlane)

	var chineese string = "\U00004e16"
	fmt.Println("-->", chineese)
}
