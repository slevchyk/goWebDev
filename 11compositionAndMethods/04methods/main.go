package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type Person struct {
	Name string
	Age  int
}

func (p Person) SomeProcessing() int {
	return 7
}

func (p Person) AgeDbl() int {
	return p.Age * 2
}

func (p Person) TakesArg(x int) int {
	return x + 1
}

func main() {

	p := Person{
		"Serhii",
		31,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}

}
