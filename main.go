package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}



func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	http.ListenAndServe(":8080", mux)
}
