package main

import (
	"fmt"
	"log"
	"net/http"
)

// handle the logic in a closure
func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, message)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/welcome", messageHandler("Welcome to writing servers with go"))
	mux.Handle("/message", messageHandler("This is a message on the message endpoint"))

	log.Println("Listening... ")
	http.ListenAndServe(":8080", mux)
}
