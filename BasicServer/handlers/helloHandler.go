package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HelloHandler struct {
	l *log.Logger
}

func NewHelloHandler(l *log.Logger) *HelloHandler {
	return &HelloHandler{l}
}

func (h *HelloHandler) ServeHttp(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusBadGateway)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}
