package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type Person struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	p1 := Person{
		"Jamed Bond",
		21,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}
}
