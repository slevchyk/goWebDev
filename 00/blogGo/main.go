package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/slevchyk/goWebDev/00/blogGo/models"
	"github.com/slevchyk/goWebDev/00/blogGo/utils"
)

var tpl *template.Template
var posts map[string]*models.Post

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	fmt.Println("Listening on port :3000")

	posts = make(map[string]*models.Post, 0)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/savePost", savePostHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delete", deleteHandler)

	http.ListenAndServe(":80", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(posts)

	err := tpl.ExecuteTemplate(w, "index", posts)
	if err != nil {
		fmt.Printf(err.Error())
		fmt.Fprintf(w, err.Error())
	}
}

func writeHandler(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "write", nil)
	if err != nil {
		fmt.Printf(err.Error())
		fmt.Fprintf(w, err.Error())
	}
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	titlte := r.FormValue("title")
	content := r.FormValue("content")

	var post *models.Post

	if id != "" {
		post = posts[id]
		post.Title = titlte
		post.Content = content
	} else {
		id := utils.GenerateId()
		post = models.NewPost(id, titlte, content)
		posts[id] = post
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	post, found := posts[id]
	if !found {
		http.NotFound(w, r)
	}

	err := tpl.ExecuteTemplate(w, "write", post)
	if err != nil {
		fmt.Printf(err.Error())
		fmt.Fprintf(w, err.Error())
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	if id == "" {
		http.NotFound(w, r)
	}

	delete(posts, id)

	http.Redirect(w, r, "/", http.StatusFound)
}
