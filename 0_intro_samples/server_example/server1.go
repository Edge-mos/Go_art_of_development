package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, rec *http.Request) {
	fmt.Println("TEST!")
	fmt.Fprintf(writer, "URL.Path = %q\n", rec.URL.Path)
}
