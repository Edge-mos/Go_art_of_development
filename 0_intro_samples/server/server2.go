package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutex sync.Mutex
var count int

func main() {
	fmt.Println("Intro")
	http.HandleFunc("/plus", mainHandler)
	http.HandleFunc("/count", counterHandler)
	http.HandleFunc("/", defaultHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	mutex.Lock()
	count++
	mutex.Unlock()
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}

func counterHandler(writer http.ResponseWriter, request *http.Request) {
	mutex.Lock()
	fmt.Fprintf(writer, "Count: %d\n", count)
	mutex.Unlock()
}

func defaultHandler(writer http.ResponseWriter, request *http.Request) {
	mutex.Lock()
	fmt.Fprintf(writer, "Default entry")
	mutex.Unlock()

}
