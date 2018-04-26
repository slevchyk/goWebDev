package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	var strFile string
	if r.Method == http.MethodPost {

		//opening file
		file, header, err := r.FormFile("FName")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		//info for debug
		fmt.Printf("\nfile: %s\nheader: %s\nerr :%s\n", file, header, err)

		//reading file
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fileOnServer, err := os.Create(filepath.Join("./usersData/", header.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer fileOnServer.Close()

		_, err = fileOnServer.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		strFile = string(bs)
	}

	err := tpl.ExecuteTemplate(w, "tpl.gohtml", strFile)
	if err != nil {
		log.Fatalln(err)
	}
}
