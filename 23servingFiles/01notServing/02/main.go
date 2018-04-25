package _1ioCopy

import (
	"net/http"
	"io"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	var body string

	body = "<h1>Hello</h1>"
	body += "<br>"
	//dosen't serve
	body += `<img src="../../rivne.png">`

	io.WriteString(w, body)
}
