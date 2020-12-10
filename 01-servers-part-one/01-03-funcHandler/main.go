package main

import (
	"fmt"
	"net/http"
	"std/log"
)

func messageHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is the message from the messageHandler")
}

func main() {
	mux := http.NewServeMux()

	// convert the messageHandler to a HandlerFunc
	mh := http.HandlerFunc(messageHandler)
	mux.Handle("/welcome", mh)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
