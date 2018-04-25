package _1

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
	body += "<img src=https://upload.wikimedia.org/wikipedia/commons/thumb/7/78/Coat_of_Arms_of_Rivne_Oblast.svg/252px-Coat_of_Arms_of_Rivne_Oblast.svg.png>"

	io.WriteString(w, body)

}
