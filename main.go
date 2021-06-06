package main

//Building Microservices with Go on youtube by Nic Jackson

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LakhanRathi92/GoBasics/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	byeHandler := handlers.NewGoodBye(l)
	sm := http.NewServeMux()

	sm.Handle("/hello", helloHandler)
	sm.Handle("/bye", byeHandler)

	s := &http.Server{
		Addr:         "127.0.0.1:9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}
