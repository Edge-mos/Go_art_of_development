package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Выводит текст каждой строки, которая появляется в стандартном вводе более одного раза, а так же количество её появлений
*/

func main() {

	counts := make(map[string]int) // это мапа
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the string(s): ")
	for scanner.Scan() {

		line := scanner.Text()
		if line == "" {
			break
		}

		counts[line]++
	}

	//игнорируем потенциальные ошибки из scanner.Err()
	for line, inputs := range counts {
		if inputs > 1 {
			fmt.Printf("Строка: %s повторяется: %d раз", line, inputs)
		}
	}
}
