package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	h1 := func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Something went wrong", http.StatusBadGateway)
			return
		}
		fmt.Fprintf(rw, "Hello %s", d)
	}

	h2 := func(w http.ResponseWriter, _ *http.Request) {
		log.Println("Good bye World")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/goodbye", h2)

	http.ListenAndServe("127.0.0.1:9090", nil)
}
