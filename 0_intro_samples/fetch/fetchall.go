package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Выполняет параллельную выборку URL и сообщает о затраченном времени и размере ответа для каждого из них
func main() {

	start := time.Now()
	channel := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, channel)
	}

	for range os.Args[1:] {
		fmt.Println(<-channel)
		//Получение из канала
	}
	fmt.Printf("%.2fs elapsed in Main \n", time.Since(start).Seconds())
}

func fetch(url string, channel chan<- string) {
	start := time.Now()

	resp, err := http.Get(url) //чтение ресурса по сети
	if err != nil {
		channel <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body) //чтение полученного ресурса
	resp.Body.Close()
	if err != nil {
		channel <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	seconds := time.Since(start).Seconds()

	channel <- fmt.Sprintf("%.2fs %7d %s", seconds, nbytes, url)
}
