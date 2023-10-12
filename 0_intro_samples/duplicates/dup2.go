package main

import (
	"bufio"
	"fmt"
	"os"
)

// Надо добавить в аргументы программы путь до файла

func main() {
	fmt.Println("Start!")

	counts := make(map[string]int)
	args := os.Args[1:]

	if len(args) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range args {

			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}

	for line, inputs := range counts {
		if inputs > 1 {
			fmt.Printf("Строка: %s повторяется: %d раз", line, inputs)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {

		line := input.Text()
		if line == "" {
			continue
		}
		counts[line]++
	}
}
