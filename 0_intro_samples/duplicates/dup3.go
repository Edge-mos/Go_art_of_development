package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {

		//чтение файла
		fileData, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		//работа с контентом
		onlyWords := strings.ReplaceAll(string(fileData), "\n", " ")
		for _, word := range strings.Split(onlyWords, " ") {
			counts[word]++
		}
	}

	//вывод результата
	for word, input := range counts {
		if input > 1 {
			fmt.Printf("Word: %q --> %d inputs\n", word, input)
		}
	}
}
