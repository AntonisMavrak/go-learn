package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func handler(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte("Welcome to Go Web Server"))
	if err != nil {
		log.Fatal(err)
		return
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		log.Fatal(err)
		return
	}

}
