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

type Regions struct {
	Name string
}

type hotel struct {
	Name   string
	Adress string
	City   string
	Zip    string
	Region Regions
}

func main() {

	Southern := Regions{"Southern"}
	Central := Regions{"Central"}
	//Northern := Regions{"Northern"}

	hotels := []hotel{
		{"Royal",
			"1st street",
			"San Diego",
			"01001",
			Southern,
		},
		{
			"Hilton",
			"2nd street",
			"Lac Vegas",
			"02001",
			Central,
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)
	if err != nil {
		log.Fatalln(err)
	}

}
