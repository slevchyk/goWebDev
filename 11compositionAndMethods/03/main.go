package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type course struct {
	Number string
	Name   string
	Units  string
}

type semestr struct {
	Term    string
	Courses []course
}

type year struct {
	Fall   semestr
	Spring semestr
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	y2018 := year{
		semestr{
			"Fall",
			[]course{
				course{"CSCI-40", "Introduction to GO", "4"},
				course{"CSCI-130", "Introduction to Web Go", "4"},
				course{"CSCI-140", "Mobile apps wuth Go", "4"},
			},
		},
		semestr{
			"Spring",
			[]course{
				course{"CSCI-40", "Advanced GO", "4"},
				course{"CSCI-130", "Advanced Web Go", "4"},
				course{"CSCI-140", "Advanced Mobile apps wuth Go", "4"},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", y2018)
	if err != nil {
		log.Fatalln(err)
	}
}
