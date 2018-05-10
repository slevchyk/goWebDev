package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "sa:sqlpass@/godata")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)

}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	_, err := io.WriteString(w, "Yeah, db connected")
	check(err)
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
