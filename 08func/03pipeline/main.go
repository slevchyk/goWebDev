package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdbl":  fdbl,
	"fsq":   fsq,
	"fsqrt": fsqrt,
}

func fdbl(f float64) float64 {
	return f * 2
}

func fsq(f float64) float64 {
	return math.Pow(f, 2)
}

func fsqrt(f float64) float64 {
	return math.Sqrt(f)
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 45.6)
	if err != nil {
		log.Fatalln(err)
	}
}
