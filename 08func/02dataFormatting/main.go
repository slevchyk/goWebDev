package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var fm = template.FuncMap{
	"fDateDMY": dayMontYear,
}

var tpl *template.Template

func dayMontYear(t time.Time) string {
	return t.Format("Monday 02 January 2006 03:04:05")
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {

	err := tpl.Execute(os.Stdout, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
