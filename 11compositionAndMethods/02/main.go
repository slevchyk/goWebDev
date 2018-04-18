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

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	p1 := DoubleZero{
		Person{
			"Jamed Bond",
			21,
		},
		true,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}
}
