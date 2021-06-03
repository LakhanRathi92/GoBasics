package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello World")
	})

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		log.Println("Good bye World")
	}

	http.HandleFunc("/goodbye", h1)

	http.ListenAndServe("127.0.0.1:9090", nil)
}
