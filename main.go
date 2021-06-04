package main

//Building Microservices with Go on youtube by Nic Jackson

import (
	"log"
	"net/http"
	"os"

	"github.com/LakhanRathi92/GoBasics/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	sm := http.NewServeMux()

	sm.Handle("/", hh)
	http.ListenAndServe("127.0.0.1:9090", sm)
}
