package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

type RootController struct {
	Tpl *template.Template
}

func NewRootController(tpl *template.Template) *RootController {
	return &RootController{
		Tpl: tpl,
	}
}

func (rc RootController) IndexHandler(w http.ResponseWriter, r *http.Request) {

	err := rc.Tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}

}
