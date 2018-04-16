package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {

	sages := map[string]string{
		"India":    "Gandi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatal(err)
	}
}
