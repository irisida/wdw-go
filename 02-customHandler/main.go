package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

func (m *messageHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	mux := http.NewServeMux()

	mh0 := &messageHandler{"I am the default, or home route"}
	mux.Handle("/", mh0)

	mh1 := &messageHandler{"A welcome message"}
	mux.Handle("/welcome", mh1)

	mh2 := &messageHandler{"the net/http package is fabulous"}
	mux.Handle("/net", mh2)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
