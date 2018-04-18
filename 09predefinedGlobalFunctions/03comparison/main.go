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

	data := struct {
		Score1 float64
		Score2 float64
		Score3 float64
		Score4 float64
	}{
		7.02,
		12.13,
		14.05,
		14.05,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
