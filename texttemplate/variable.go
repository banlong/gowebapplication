package main

import (
	"log"
	"os"
	"text/template"
)

type Note struct {
	Title string
	Description string
}
const tmpl =
   `Note - Title: {{.Title}}, Description: {{.Description}}
	{{ $note := "Sample Note"}}
	{{ $note }}
	{{ $a := 5}}
	{{ $b := 5}}
	{{ if eq $a $b }} a and b are equal{{ end }}
	{{ if eq $a $b }} {{$note | println }} {{ end }}
	`

func main() {
	//Create an instance of Note struct
	note := Note{"text/templates", "Template generates textual output"}
	//create a new template with a name
	t := template.New("note")
	//parse some content and generate a template
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	//Applies a parsed template to the data of Note object
	if err := t.Execute(os.Stdout, note); err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}


