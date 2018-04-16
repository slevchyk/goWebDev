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

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "love all",
	}

	sages := []sage{buddha, gandhi, jesus}

	bmw := car{
		Manufacturer: "BMW",
		Model:        "740",
		Doors:        4,
	}

	porche := car{
		Manufacturer: "Porsche",
		Model:        "911",
		Doors:        2,
	}

	vw := car{
		Manufacturer: "VW",
		Model: "Golf",
		Doors: 5,
	}

	cars := []car{bmw, porche, vw}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
