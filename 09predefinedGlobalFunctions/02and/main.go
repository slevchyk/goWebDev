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

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {

	u0 := user{
		"Buddha",
		"The belief of no beliefs",
		false,
	}

	u1 := user{
		"Gandhi",
		"Be the change",
		true,
	}

	u2 := user{
		"",
		"Nothing is nothing",
		true,
	}

	data := []user{u0, u1, u2}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
