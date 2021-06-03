package handlers

import (
	"log"	"net/http"
)


type HelloHandler struct {
	l *log.Logger
}

func newHelloHandler(l *log.Logger) {
	&HelloHandler{l}
}

func (h* HelloHandler) serveHttp(rw http.ResponseWriter, r *http.Request){
	h.l.Println("Hello World");
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusBadGateway)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}
