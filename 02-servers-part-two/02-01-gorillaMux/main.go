package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"created_on"`
}

// store for the Notes collection
var noteStore = make(map[string]Note)

var id int = 0

// HTTP Post - /api/notes
func PostNoteHandler(w http.ResponseWriter, req *http.Request) {
	var note Note

	// decode the incoming json
	err := json.NewDecoder(req.Body).Decode(&note)
	if err != nil {
		panic(err)
	}

	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note

	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// HTTP Get - /api/notes
func GetNoteHandler(w http.ResponseWriter, req *http.Request) {
	var notes []Note

	for _, v := range noteStore {
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// HTTP Put - /api/notes/{id}
func PutNoteHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	vars := mux.Vars(req)
	k := vars["id"]
	var noteToUpd Note

	// decode the incoming json
	err = json.NewDecoder(req.Body).Decode(&noteToUpd)
	if err != nil {
		panic(err)
	}

	if note, ok := noteStore[k]; ok {
		noteToUpd.CreatedOn = note.CreatedOn

		// delete existing note and add updated item at the same index
		delete(noteStore, k)
		noteStore[k] = noteToUpd
	} else {
		log.Printf("Could not find key of note %s tp delete", k)
	}
	w.WriteHeader(http.StatusNoContent)

}

// HTTP Delete - /api/notes/{id}
func DeleteNoteHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	k := vars["id"]

	// remove from store
	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		log.Printf("Could not find the key of note %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	server.ListenAndServe()
}
