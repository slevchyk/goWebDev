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

	err := tpl.Execute(os.Stdout, "Serhii")
	if err != nil {
		log.Fatal(err)
	}

}
