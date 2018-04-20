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

type Page struct {
	Title   string
	Heading string
	Input   string
}

func main() {

	home := Page{
		"New page",
		"It's a heading",
		`<script>alert("wohoo");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}
