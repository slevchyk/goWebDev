package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
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
		Model:        "Golf",
		Doors:        5,
	}

	cars := []car{bmw, porche, vw}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
