package main

import (
	"log"
	"os"
	"text/template"
)

// Note - struct used to hold fileds of a note
type Note struct {
	Title       string
	Description string
}

const tmpl = `Notes are: 
{{range .}}
- Title: {{.Title}}, Description: {{.Description}}
{{end}}`

func main() {

	// create an slice of the note struct
	// sample uses two notes for demo purposes
	notes := []Note{
		{"text/templates", "Templates generate the textual output"},
		{"html/templates", "Templates generate the HTML output"},
	}

	// create a new template
	t := template.New("notes")

	// parse content and generate template
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	// applies parsed template to the data of note object
	if err := t.Execute(os.Stdout, notes); err != nil {
		log.Fatal("Execute: ", err)
		return
	}

}
