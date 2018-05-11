package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("HTTPS server\n"))

}
