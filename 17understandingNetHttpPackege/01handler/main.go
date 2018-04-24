package main

import (
	"fmt"
	"net/http"
)

type someType string

func (st someType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, " - you can do any you want")
}

func main() {

}
