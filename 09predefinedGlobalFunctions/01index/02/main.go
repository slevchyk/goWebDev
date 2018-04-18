package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	ss := []string{"zero", "one", "two", "three", "four", "five"}

	data := struct {
		Words []string
		Lname string
	}{
		ss,
		"Levchyk",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
