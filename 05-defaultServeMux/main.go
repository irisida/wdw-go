package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHAndler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Some message we want ro reach")
}

func main() {
	http.HandleFunc("/reach", messageHAndler)

	log.Println("Listening...")

	// note we're using nil as the second
	// parameter where we would normally
	// add the variable of the mux we have
	// created. A nil demotes the use of
	// the default built-in mux
	http.ListenAndServe(":8080", nil)
}
