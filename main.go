package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/forms", formHandler)
	http.HandleFunc("/index", indexHandler)

	fmt.Println("Listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/index" {
		http.Error(writer, "404 page not found", http.StatusNotFound)
	}
	if request.Method != "GET" {
		http.Error(writer, "405 method not allowed", http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(writer, "Hello World!")
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v", err)
	}
	fmt.Fprintf(writer, "POST request successful")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(writer, "name=%s\n, address=%s\n", name, address)
}
