package types

import (
	"fmt"
	"unicode/utf8"
)

//Строки - неизменяемые последовательности байтоа

func StringsEx() {
	fmt.Println("Строки практика")

	str := "hello, world"

	fmt.Printf("длина в байтах: %d \n", len(str))
	fmt.Printf("первый байт:%d (%c), пятый байт: %d(%c) \n", str[0], str[0], str[7], str[7])

	bytes := str[:]
	fmt.Println(bytes) //выведет визуально как строку
	fmt.Println("Сравним слайс байтов и строку", bytes == "hello, world")

	for _, byteVal := range bytes {
		fmt.Printf("%c \n", byteVal)
	}

	fmt.Println("--------взятие подстроки------")
	fmt.Println(str[:5])
	fmt.Println(str[7:])

}

func HasPrefixInStr(str, prefix string) bool {
	return len(str) >= len(prefix) && str[:len(prefix)] == prefix
}

func RunesInStringEx() {
	fmt.Println("\n==========Runes in string=========")

	str := "Hello, " + "\U00004e16\U0000754c"
	fmt.Println("Исходная строка str: ", str)
	fmt.Println("длина str в байтах = ", len(str))
	fmt.Println("длина str в символах = ", utf8.RuneCountInString(str))

	for i := 0; i < len(str); {

		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("index: %[1]d\t intVal: %[2]d\t Rune: %[2]c \n", i, r)
		i += size
	}

	fmt.Println("-------ошибка с не ASCII символами")
	for i := 0; i < len(str); i++ {
		fmt.Printf("index: %[1]d\t intVal: %[2]d\t Rune: %[2]c \n", i, str[i])
	}

	fmt.Println("Если пройтись по диапазону в строке, то екодирование происходит неявно")
	runeCount := 0
	for i, rune := range str {
		fmt.Printf("%d\t%q\t%d\n", i, rune, rune)
		runeCount++
	}
	fmt.Println("всего рун = ", runeCount)

	r := []rune(str) //[72 101 108 108 111 44 32 19990 30028]
	fmt.Println(r)
	fmt.Printf("%x\n", r)

}
