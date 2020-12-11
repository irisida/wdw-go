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

const tmpl = `Note - Title: {{.Title}}, Description: {{.Description}}`

func main() {

	// create an instance of the note struct
	note := Note{"text/templates", "Templates generate the textual output"}

	// create a new template
	t := template.New("note")

	// parse content and generate template
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	// applies parsed template to the data of note object
	if err := t.Execute(os.Stdout, note); err != nil {
		log.Fatal("Execute: ", err)
		return
	}

}
