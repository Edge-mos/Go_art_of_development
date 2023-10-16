package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//Выводит ответ на запрос по заданному URL

func main() {

	for _, url := range os.Args[1:] {

		//получаем ответ по сети
		resp, err := http.Get(url)

		//ошибка, выход из программы
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		//извлечение тела ответа
		bodyData, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}

		//вывод тела ответа(идет автокастинг через формат)
		fmt.Printf("%s", bodyData)
	}
}
